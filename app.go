package main

import (
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!\n"))
}

func main() {
	http.HandleFunc("/", Handler)
	log.Println("Start HTTP server on port 8881")
	log.Fatal(http.ListenAndServe(":8881", nil))
}
