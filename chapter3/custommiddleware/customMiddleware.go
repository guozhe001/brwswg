package main

import (
	"fmt"
	"net/http"
)

func middleware(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("executing middleware before request phase!")
		handlerFunc.ServeHTTP(w, r)
		fmt.Println("executing middleware after response phase!")
	}
}

func mainLogic(w http.ResponseWriter, r *http.Request) {
	fmt.Println("executing mainHandler...")
	w.Write([]byte("OK"))
}

func main() {
	handlerFunc := http.HandlerFunc(mainLogic)
	http.HandleFunc("/", middleware(handlerFunc))
	http.ListenAndServe(":8088", nil)
}
