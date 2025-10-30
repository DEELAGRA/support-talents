package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct {
	ID   int
	Name string
}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Привет ты запустил сервер твой ID: %v\n", h.ID)
}

func main() {

	fmt.Println("Сервер зарущен на http://localhost:8080")
	http.ListenAndServe(":8080", MyHandler{})
}
