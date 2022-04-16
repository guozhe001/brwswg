package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category is: %v\n", vars["category"])
	fmt.Fprintf(w, "ID is: %v\n", vars["id"])
}

func main() {
	r := mux.NewRouter()
	s := r.PathPrefix("/articles").Subrouter()
	s.HandleFunc("", QueryHandler)
	s.Queries("id", "category")
	s.HandleFunc("/{category}/{id:[0-9]+}", ArticleHandler)
	s.HandleFunc("/{id:[0-9]+}/settings", settingsHandler)
	s.HandleFunc("/{id:[0-9]+}/details", detailsHandler)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("/Users/guozhe/code/my-project/brwswg//static"))))
	server := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8088",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}

func QueryHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Got parameter Category is: %v\n", query["category"])
	fmt.Fprintf(w, "Got parameter ID is: %v\n", query["id"])
}

func detailsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ID is: %v\n", vars["id"])
	fmt.Fprintf(w, "this is detailsHandler")
}

func settingsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ID is: %v\n", vars["id"])
	fmt.Fprintf(w, "this is settingsHandler")
}
