package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct {
	Name  string
	Count int
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Count++
	response := fmt.Sprintf("Привет от %s! Это ваш запрос № %d\n", h.Name, h.Count)
	fmt.Fprint(w, response)
}

func main() {

	fmt.Println("Сервер зарущен на http://localhost:8080")
	http.ListenAndServe(":8080", &MyHandler{Name: "HTTP Server"})
}
