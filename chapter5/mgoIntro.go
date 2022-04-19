package main

//
//import (
//	"context"
//	"fmt"
//	"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/mongo"
//	"go.mongodb.org/mongo-driver/mongo/options"
//	"log"
//	"time"
//)
//
//type BoxOffice struct {
//	Budget uint64 `bson:"budget"`
//	Gross  uint64 `bson:"gross"`
//}
//
//type Movie struct {
//	Name      string   `bson:"name"`
//	Year      string   `bson:"year"`
//	Directors []string `bson:"directors"`
//	Writers   []string `bson:"writers"`
//	BoxOffice `bson:"boxOffice"`
//}
//
//func main() {
//	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancelFunc()
//	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:123456@localhost:27017/"))
//	if err != nil {
//		panic(any(err))
//	}
//	c := client.Database("appdb").Collection("movies")
//
//	// Create
//	movie := &Movie{Name: "The Dark Knight",
//		Year:      "2008",
//		Directors: []string{"Christopher Nolan"},
//		Writers:   []string{"Jonathan Nolan", "Christopher Nolan"},
//		BoxOffice: BoxOffice{
//			Budget: 185000000,
//			Gross:  533316061,
//		},
//	}
//	c.InsertOne(ctx, movie)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// query
//	result := Movie{}
//	err = c.FindOne(ctx, bson.M{"boxOffice.budget": bson.M{"$gt": 150000000}}).Decode(&result)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("movie:", result.Name)
//
//}
