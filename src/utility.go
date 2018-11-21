package main

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

func dbInit() *sqlx.DB {
	username := os.Getenv("OGRE_DB_USERNAME")
	password := os.Getenv("OGRE_DB_PASSWORD")
	host := os.Getenv("OGRE_DB_HOSTNAME")
	dbname := os.Getenv("OGRE_DB_NAME")
	db := dbConnect(username, password, host, dbname)

	return db
}

func dbConnect(username string, password string, host string, dbname string) *sqlx.DB {
	db, err := sqlx.Connect("mysql", username+":"+password+"@"+host+"/"+dbname)
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
