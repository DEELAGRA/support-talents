package main

import "fmt"

func main() {
	fmt.Println("Привет из paskal!")

	sliceOfslice := [][]int{}
	num := 5
	for i := 0; i < num; i++ {
		arr := make([]int, i+1)
		arr[0] = 1
		arr[i] = 1
		if i > 0 {
			for j := 1; j < i; j++ {
				arr[j] = sliceOfslice[i-1][j-1] + sliceOfslice[i-1][j]
			}
		}

		sliceOfslice = append(sliceOfslice, arr)

	}
	fmt.Println(sliceOfslice)

}
