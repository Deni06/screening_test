package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

type ListMuseum struct {
	Result []struct{
		MuseumID string `json:"museum_id"`
		KodePengelolaan string `json:"kode_pengelolaan"`
		Nama string `json:"nama"`
		Sdm string `json:"sdm"`
		AlamatJalan string `json:"alamat_jalan"`
		DesaKelurahan string `json:"desa_kelurahan"`
		Kecamatan string `json:"kecamatan"`
		KabupatenKota string `json:"kabupaten_kota"`
		Propinsi string `json:"propinsi"`
		Lintang string `json:"lintang"`
		Bujur string `json:"bujur"`
		Koleksi string `json:"koleksi"`
		SumberDana string `json:"sumber_dana"`
		Pengelola string `json:"pengelola"`
		Tipe string `json:"tipe"`
		Standar string `json:"standar"`
		TahunBerdiri string `json:"tahun_berdiri"`
		Bangunan string `json:"bangunan"`
		LuasTanah string `json:"luas_tanah"`
		StatusKepemilikan string `json:"status_kepemilikan"`
} `json:"data"`
}

var wg sync.WaitGroup

func main() {

	concurrentLimit := flag.Int("concurrent_limit", 2, "Limit Concurrency")
	output := flag.String("output", "", "Output file excel")

	flag.Parse()

	url := "http://jendela.data.kemdikbud.go.id/api/index.php/CcariMuseum/searchGET?nama=museum"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	body := bytes.TrimPrefix(bodyBytes, []byte("\xef\xbb\xbf"))

	var record ListMuseum

	err = json.Unmarshal([]byte(body), &record)
	if err != nil {
		log.Fatal(err)
	}

	queue := make(chan string, *concurrentLimit)

	dataMuseum := make(map[string][][]string)

	for _, data := range record.Result{
		dataMuseum[data.KabupatenKota] = append(dataMuseum[data.KabupatenKota], []string{data.MuseumID, data.KodePengelolaan, data.Nama, data.Sdm, data.AlamatJalan,
			data.DesaKelurahan, data.Kecamatan, data.KabupatenKota, data.Propinsi, data.Lintang, data.Bujur, data.Koleksi,
			data.SumberDana, data.Pengelola, data.Tipe, data. Standar, data.TahunBerdiri, data.Bangunan, data.LuasTanah,
			data.StatusKepemilikan})
	}
		for name, listExcel := range dataMuseum {
			wg.Add(1)
			go getData(queue, listExcel, *output, name)
		}

	go func() {
		wg.Wait()
		close(queue)
	}()

	for data := range queue {
		fmt.Println(data)
	}
}

func getData(c chan<- string, listExcel [][]string, output string, name string){
	defer wg.Done()
	err := csvExport(listExcel, filepath.Join(output, name+".csv"))
	if err != nil {
		log.Fatal(err)
	}
	c <- name+".csv Generated"
}

func csvExport(data [][]string, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		if err := writer.Write(value); err != nil {
			return err // let's return errors if necessary, rather than having a one-size-fits-all error handler
		}
	}
	return nil
}