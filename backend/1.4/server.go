package main

import (
	"fmt"
	"net/http"
	"time"
)

func serverTime() {
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/", helloHandler)
	fmt.Println("Запуск сервера на http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	if r.URL.Path == "/time" {
		fmt.Fprintf(w, "Сейчас %v ты ещё успеваешь?", now)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprintln(w, "Ты попал на мой сервер")
		fmt.Fprintln(w, "Если перейде по ссылке http://localhost:8080/time, то узнаешь который час.")
	}
}
