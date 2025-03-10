package main

import (
	"fmt"
	"net/http"
)

func handlerHello(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Write([]byte("Hello World"))
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func main() {
	http.HandleFunc("/", handlerHello)
	fmt.Println("Server running of :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
