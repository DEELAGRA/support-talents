package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Была запущена функция -файл-")
	file()

	time.Sleep(3 * time.Second)

	fmt.Println("Была запущена функция -сервера-")

	target := 12
	arr := []int{1, 7, 2, 9, 12, 45, 32, 77, 81, 10, 99}
	fmt.Println("Массив до сортировки:", arr)
	bublesort(arr)
	time.Sleep(3 * time.Second)
	fmt.Println("Массив после сортировки:", arr)

	fmt.Println("Число под индексом", binarSearch(arr, target))
}
