package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	newMux := http.NewServeMux()
	newMux.HandleFunc("/randomFloat", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Your random number is: %f", rand.Float64())
	})
	newMux.HandleFunc("/randomInt", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, strconv.Itoa(rand.Intn(100)))
	})
	http.ListenAndServe(":8088", newMux)
}
