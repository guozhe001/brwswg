package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

type Book struct {
	id     int
	name   string
	author string
}

func main() {
	db_file := "/Users/guozhe/code/my-project/brwswg/chapter4/books.db"
	os.Remove(db_file)
	db, err := sql.Open("sqlite3", db_file)
	if err != nil {
		log.Println(err)
	}
	// create
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS books (id INTEGER PRIMARY KEY, isbn INTEGER, author VARCHAR(64), name VARCHAR(64) NULL)")
	if err != nil {
		log.Println("error in creating table")
	} else {
		log.Println("successfully created table books!")
	}
	statement.Exec()

	// insert
	statement, _ = db.Prepare("INSERT INTO books (name, author, isbn) VALUES (?, ?, ?)")
	statement.Exec("A table of two cities", "C D", 12332)
	log.Println("inserted the book into database!")

	// query
	rows, _ := db.Query("SELECT id, name, author FROM books")
	if rows != nil {
		var tempBook Book
		for rows.Next() {
			rows.Scan(&tempBook.id, &tempBook.name, &tempBook.author)
			log.Printf("ID:%d, Book:%s, Author:%s\n", tempBook.id, tempBook.name, tempBook.author)
		}
	}

	// update
	statement, _ = db.Prepare("update books set name = ? where id= ?")
	statement.Exec("The Table Of Two Cities", 1)
	log.Println("Successfully updated the book in database")

	// Delete
	statement, _ = db.Prepare("delete from books where id = ?")
	statement.Exec(1)
	log.Println("successfully delete the book in database")
}
