package main

import (
	"golangweb/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/product", handler.ProductHandler)

	log.Println("starting web on port 8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
