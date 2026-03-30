package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "Hello, World!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Println("Listening for requests at http://localhost:8000/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
