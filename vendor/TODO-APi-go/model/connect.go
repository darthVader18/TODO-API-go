package model

import (
	"database/sql"
	"fmt"
	"log"
)

var con *sql.DB

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "postgres://jyyjzaznhrqazf:7dffe2c89bcc2dd572cf4f246fcf11fad6d1c4ee71fad15bede2261df033aa31@ec2-54-175-117-212.compute-1.amazonaws.com:5432/d8kvaiv71qft39")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")
	con = db
	return db
}
