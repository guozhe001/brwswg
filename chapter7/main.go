package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/guozhe001/brwswg/chapter7/base62"
	models "github.com/guozhe001/brwswg/chapter7/src"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type DBClient struct {
	db *sql.DB
}

type Record struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}

func (d *DBClient) GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	var url string
	vars := mux.Vars(r)
	id := base62.ToBase10(vars["encoded_string"])
	err := d.db.QueryRow("select url from web_url where id = $1", id).Scan(&url)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		responseMap := map[string]interface{}{"url": url}
		response, _ := json.Marshal(responseMap)
		w.Write(response)
	}
}

func (d *DBClient) GenerateShortURL(w http.ResponseWriter, r *http.Request) {
	var id int
	var record Record
	postBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(postBody, &record)
	err := d.db.QueryRow("insert into web_url(url) values($1) returning id", record.URL).Scan(&id)
	responseMap := map[string]interface{}{"encoded_string": base62.ToBase62(id)}
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(responseMap)
		w.Write(response)
	}

}

func main() {
	db, err := models.InitDB()
	if err != nil {
		panic(any(err))
	}
	dbClient := &DBClient{db: db}
	defer db.Close()
	r := mux.NewRouter()
	r.HandleFunc("/v1/short/{encoded_string:[a-zA-Z0-9]*}", dbClient.GetOriginalURL).Methods(http.MethodGet)
	r.HandleFunc("/v1/short", dbClient.GenerateShortURL).Methods(http.MethodPost)
	s := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8088",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}
