package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", MyServer)
	log.Fatal(http.ListenAndServe(":8088", nil))
}

func MyServer(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Hello world!")
	//fmt.Fprintf(writer, "Hello world!")
}
