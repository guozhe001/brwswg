package gorillamiddleware

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func mainLogic(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing request!")
	fmt.Fprintf(w, "OK")
	log.Println("Finished processing request!")
}

func Main() {
	r := mux.NewRouter()
	r.HandleFunc("/", mainLogic)
	handler := handlers.LoggingHandler(os.Stdout, r)
	http.ListenAndServe(":8088", handler)
}
