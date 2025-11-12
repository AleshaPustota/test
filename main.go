package main

import (
	"net/http"
	"prac1/handlers"
)

func main() {
	http.HandleFunc("/user/recive", handlers.HandlerRecive)
	http.ListenAndServe(":8080", nil)

}
