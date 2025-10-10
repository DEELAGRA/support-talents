package main

import (
	"fmt"
)

func reverseint() {

	x := -125987435
	original := x
	reverse := 0
	if x > 0 {
		for i := 1; x > 0; i++ {
			reverse = reverse*10 + x%10

			x /= 10
		}
	} else {
		for i := 1; x < 0; i++ {
			reverse = reverse*10 + x%10

			x /= 10
		}

	}
	fmt.Printf("Reverse Int:\nИзначально было %d, стало %d\n\n", original, reverse)
}
