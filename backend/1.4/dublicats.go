package main

import "fmt"

func dublicats(arr []int) bool {

	bubblesort(arr)

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			fmt.Printf("i - %v, j - %v\n", arr[i], arr[j])
			if arr[j] == arr[i] {

				return true
			}
		}

	}
	return false
}
