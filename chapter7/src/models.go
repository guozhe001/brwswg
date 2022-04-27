package models

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func InitDB() (*sql.DB, error) {
	var err error
	db, err := sql.Open("postgres", "postgresql://postgres:123456@192.168.31.152:5432/mydb?sslmode=disable")
	if err != nil {
		return nil, err
	} else {
		stmt, err := db.Prepare("create table web_url(id serial primary key, url text not null);")
		if err != nil {
			log.Println(err)
			return nil, err
		}
		res, err := stmt.Exec()
		log.Println(res)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return db, nil
	}

}
