package main

import (
	models "github.com/guozhe001/brwswg/chapter7/src"
	"log"
)

func main() {
	db, err := models.InitDB()
	if err != nil {
		log.Println(db)
	}
}
