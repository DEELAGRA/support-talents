package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Только метод POST", http.StatusMethodNotAllowed)
		return
	}
	var text map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&text); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if text == nil {
		http.Error(w, "Отправь тело json", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(text)

}

func main() {

	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", &MyHandler{})
}
