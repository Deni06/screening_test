package util

import "fmt"

func CreateChart(charts []int){
	maxValue := int(0)
	for i := 0; i < len(charts); i++ {
		if maxValue < charts[i] {
			maxValue = charts[i]
		}
	}

	for k := maxValue; k >= 0; k-- {
		for j := 0; j < len(charts); j++ {
			if k == 0 {
				fmt.Print(fmt.Sprintf("%v ",charts[j]))
			}else if k <= charts[j] {
				fmt.Print("| ")
			}else{
				fmt.Print("  ")
			}
		}
		fmt.Println("")
	}
}

func InsertionSort(charts []int) {
	CreateChart(charts)
	fmt.Println("")
	for i := 1; i < len(charts); i++ {
		j := i
		for j > 0 {
			if charts[j-1] > charts[j] {
				charts[j-1], charts[j] = charts[j], charts[j-1]
				CreateChart(charts)
				fmt.Println("")
			}
			j -= 1
		}
	}
}

func ReverseSort(charts []int) {
	CreateChart(charts)
	fmt.Println("")
	for i := 1; i < len(charts); i++ {
		j := i
		for j > 0 {
			if charts[j-1] < charts[j] {
				charts[j-1], charts[j] = charts[j], charts[j-1]
				CreateChart(charts)
				fmt.Println("")
			}
			j -= 1
		}
	}
}
