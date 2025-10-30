package main

import "fmt"

func qucksort(slice []int, low, hight int) {
	if low >= hight {
		return
	}
	pivotIndex := partition(slice, low, hight)
	qucksort(slice, low, pivotIndex-1)
	qucksort(slice, pivotIndex+1, hight)
}

func partition(slice []int, low, hight int) int {
	pivot := slice[hight]
	i := low - 1

	for j := low; j < hight; j++ {
		if slice[j] <= pivot {

			i++
			slice[i], slice[j] = slice[j], slice[i]

		}
	}
	slice[i+1], slice[hight] = slice[hight], slice[i+1]
	return i + 1
}
func QuickSort(slice []int) {
	if len(slice) > 1 {
		qucksort(slice, 0, len(slice)-1)
	}
}

func main() {
	data := []int{64, 34, 25, 12, 22, 11, 1}

	fmt.Println("До сортировки:", data)

	QuickSort(data)

	fmt.Println("После сортировки:", data)
}
