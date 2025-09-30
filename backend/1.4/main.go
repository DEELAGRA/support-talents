package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Была запущена функция -файл-")
	file()

	time.After(1000)

	fmt.Println("Была запущена функция -сервера-")
	serverTime()
}
