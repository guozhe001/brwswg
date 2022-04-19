package main

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type DB struct {
	client     *mongo.Client
	collection *mongo.Collection
	ctx        context.Context
}

var db *DB

type Movie struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name,omitempty"`
	Year      string             `json:"year" bson:"year"`
	Directors []string           `json:"directors" bson:"directors"`
	Writers   []string           `json:"writers" bson:"writers"`
	BoxOffice `json:"boxOffice" bson:"boxOffice"`
}

type BoxOffice struct {
	Budget uint64 `json:"budget" bson:"budget"`
	Gross  uint64 `json:"gross" bson:"gross"`
}

func (d *DB) GetMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	var movie Movie
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
	}
	err = d.collection.FindOne(d.ctx, bson.M{"_id": id}).Decode(&movie)
	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(movie)
		w.Write(response)
	}

}

func (d *DB) CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	postBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(postBody, &movie)
	result, err := d.collection.InsertOne(d.ctx, movie)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		// 强制类型转换
		movie.ID = result.InsertedID.(primitive.ObjectID)
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(movie)
		w.Write(response)
	}

}

func main() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:123456@localhost:27017/"))
	if err != nil {
		panic(any(err))
	}
	defer client.Disconnect(ctx)
	c := client.Database("appdb").Collection("movies")
	db = &DB{client: client, ctx: ctx, collection: c}
	r := mux.NewRouter()
	r.HandleFunc("/v1/movies/{id:[a-zA-Z0-9]*}", db.GetMovie).Methods("GET")
	r.HandleFunc("/v1/movies", db.CreateMovie).Methods(http.MethodPost)
	s := &http.Server{
		Handler:      r,
		Addr:         ":8088",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	log.Fatal(s.ListenAndServe())

}
