package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	source := "./"
	target := "./"

	if len(os.Args) > 1{
		source = os.Args[1]
	}
	if len(os.Args) > 2{
		target = os.Args[2]
	}

	compare(source, target)
}

func compare(source, target string)  {

	sourceList := make([]string, 0)
	targetList := make([]string, 0)

	err := filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if path != source && !info.IsDir(){

			source = filepath.Join(source)

			if source == "" || source == "."{
				sourceList = append(sourceList, path)
			}else{
				splitName := strings.Split(path, source+"\\")
				if len(splitName) > 1{
					sourceList = append(sourceList, splitName[1])
				}
			}
		}
		return err
	})

	if err != nil {
		log.Fatal(err)
	}

	listFileDelete := map[string]int{}

	err = filepath.Walk(target, func(path string, info os.FileInfo, err error) error {
		if path != target && !info.IsDir(){

			target = filepath.Join(target)

			if target == "" || target == "."{
				listFileDelete[path] = 0
				targetList = append(targetList, path)
			}else{
				splitName := strings.Split(path, target+"\\")
				if len(splitName) > 1{
					listFileDelete[splitName[1]] = 0
					targetList = append(targetList, splitName[1])
				}
			}
		}
		return err
	})

	if err != nil {
		log.Fatal(err)
	}

	for _,s := range sourceList{
		isSame := false
		for _, t := range targetList {
			if s == t {
				listFileDelete[t] = 1
				isSame = true

				fileSameCheck := make([]string, 0)
				indexScan := 0

				fileSource, err := os.Open(filepath.Join(source,s))
				if err != nil {
					log.Fatal(err)
				}

				scannerSource := bufio.NewScanner(fileSource)

				for scannerSource.Scan() {
					fileSameCheck = append(fileSameCheck, scannerSource.Text())
				}

				fileSource.Close()

				fileTarget, err := os.Open(filepath.Join(target,s))
				if err != nil {
					log.Fatal(err)
				}

				scannerTarget := bufio.NewScanner(fileTarget)

				for scannerTarget.Scan() {
					if indexScan == len(fileSameCheck){
						fmt.Println(strings.Replace(s, string(filepath.Separator), "/", -1)+" MODIFIED")
						break
					}
					if scannerTarget.Text() != fileSameCheck[indexScan] {
						fmt.Println(strings.Replace(s, string(filepath.Separator), "/", -1)+" MODIFIED")
						break
					}
					indexScan++
				}

				fileTarget.Close()

					break
			}
		}
		if !isSame {
			fmt.Println(strings.Replace(s, string(filepath.Separator), "/", -1)+" NEW")
		}
	}

	for key,lfd := range listFileDelete{
		if lfd == 0 {
			fmt.Println(strings.Replace(key, string(filepath.Separator), "/", -1)+" DELETED")
		}
	}
}