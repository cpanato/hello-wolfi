package main

import (
	"log"
	"net/http"
)

type HelloHandler struct{}

func (handler HelloHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello Wolfi!\n"))
}

func main() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: HelloHandler{},
	}

	log.Println("Server starting...")
	log.Fatal(server.ListenAndServe())
}
