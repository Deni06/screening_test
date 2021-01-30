#### Solution 3
##### Concurrency Task Worker

Contoh compile :
```
Windows:
grab.exe -concurrent_limit=2 -output=/home/bayu/museum

Linux:
grab -concurrent_limit=2 -output=/home/bayu/museum
```

##### Penjelasan Program

- Pertama, kita siapkan variabel WaitGroup, Waitgroup dapat dianggap sebagai penghitung global yang dapat kita gunakan untuk menunggu semua goroutine selesai
- Lalu kita harus mengget API dari url http://jendela.data.kemdikbud.go.id/api/index.php/CcariMuseum/searchGET?nama=museum menggunakan method GET
- Lalu kita akan mengkonversi response body yang kita dapatkan ke byte[] menggunakan **ioutil.ReadAll()** 
- Lalu kita petakan hasil konversi tadi ke struct yang disini bernama **"ListMuseum"**
- Siapkan satu variabel bertipe **map[string][][]string**, variabel ini digunakan saat kita akan melakukan generate ke csv
- Ketika kita melakukan perulangan, kita dapat mengelompokkan data tersebut berdasarkan kabupaten/kota dengan cara kita mengpush data tersebut ke variabel bertipe map tadi dengan dengan index = value kabupaten/kota
- Siapkan satu variabel yang akan digunakan sebagai channel untuk goroutine dan set limit concurrency nya sesuai dengan yang didapat dari parameter
- Lalu kita lakukan perulangan dari variabel map yang tadi kita buat, dan menambahkan statement **wg.Add(1)** yang digunakan untuk menunjukkan bahwa kita sedang membuat goroutine. 
- Lalu melakukan panggilan ke fungsi **getData(c chan<- string, listExcel [][]string, output string, name string)** menggunakan goroutine sehingga prosesnya akan dijalankan secara asynchronus
- Lalu fungsi **getData(c chan<- string, listExcel [][]string, output string, name string)** akan memanggil fungsi lain yaitu fungsi **csvExport(data [][]string, fileName string)**. Difungsi ini lah proses generate csv berlangsung dimana data csv ini akan disimpan didirectory sesuai dengan parameter yang dikirim.
- Setelah proses generate csv selesai, kita akan menjalankan **wg.Done()** difungsi **getData(c chan<- string, listExcel [][]string, output string, name string)** yang menunjukkan bahwa goroutine untuk proses ini telah selesai
- Setelah semua perulangan dilakukan kita menutup proses goroutine dengan statement **wg.Wait**. Statement **wg.Wait()** bersifat blocking, proses eksekusi program tidak akan diteruskan ke baris selanjutnya, sebelum semua goroutine selesai. Terakhir, tutup channel yang kita pakai dengan statement **close(queue)**