package main

import "fmt"

func movezeros(arr []int) {
	fmt.Printf("Был такой: %v\n", arr)
	for i, num := range arr {

		if num == 0 {
			arr = append(arr[:i], arr[i+1:]...)
			arr = append(arr, num)
		}
	}
	fmt.Printf("Move Zeros:\nСтал таким: %d\n\n", arr)
}
