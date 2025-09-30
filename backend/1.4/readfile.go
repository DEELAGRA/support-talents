package main

import (
	"fmt"
	"log"
	"os"
)

func file() {
	err := os.WriteFile("hellofile.txt", []byte("Привет программист!"), 0600)
	if err != nil {
		log.Fatal(err)
	}

	data, err := os.ReadFile("hellofile.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
	err = os.Remove("hellofile.txt")
	if err != nil {
		log.Fatal(err)
	}
}
