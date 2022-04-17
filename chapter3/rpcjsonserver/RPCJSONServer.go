package main

import (
	jsonparse "encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Args struct {
	Id string
}

type Book struct {
	Id     string "json:string,omitempty"
	Name   string "json:string,omitempty"
	Author string "json:string,omitempty"
}

type JSONServer struct {
}

func (j *JSONServer) GiveBookDetail(r *http.Request, args *Args, reply *Book) error {
	var books []Book
	raw, err := ioutil.ReadFile("./books.json")
	if err != nil {
		log.Println("ReadFile error:", err)
		os.Exit(1)
	}
	err = jsonparse.Unmarshal(raw, &books)
	if err != nil {
		log.Println("json Unmarshal error:", err)
		os.Exit(1)
	}
	for _, book := range books {
		if book.Id == args.Id {
			*reply = book
			break
		}
	}
	return nil
}

func main() {
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterService(new(JSONServer), "")
	r := mux.NewRouter()
	r.Handle("/rpc", s)
	http.ListenAndServe(":1234", r)
}
