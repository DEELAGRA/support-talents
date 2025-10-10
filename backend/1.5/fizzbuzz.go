package main

import (
	"fmt"
	"strconv"
)

func fizzbuzz(n int) {
	arr := []string{}
	for i := 1; i < n+1; i++ {
		x := strconv.Itoa(i)
		if i%3 == 0 {
			arr = append(arr, "fizz")
		} else if i%5 == 0 {
			arr = append(arr, "bizz")
		} else {
			arr = append(arr, x)
		}

	}
	fmt.Printf("Fizz Bizz:\nБыло число: %v \nМассив сейчас такой: %v\n\n", n, arr)

}
