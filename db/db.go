package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Create a global variable so that other applications can use this same databse or interact with this database
var DB *sql.DB

//creating a function to make sure we have a connection and we can work with a database

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Could not connect to database " + err.Error())
	}

	//to manage connections, specifying how many connections it needs
	DB.SetMaxOpenConns(10)
	//how many connections you want to leave open, if no one is using the connections at the moment we can set max
	DB.SetMaxIdleConns(5)

	createTables(DB) //call function to execute create table
}

//create tables for the event

func createTables(db *sql.DB) {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	location TEXT NOT NULL,
	dateTime DATETIME NOT NULL,
	user_id INTEGER
	)`

	//To execute a query statement
	//_, err := DB.Exec(createEventsTable) // store event table in a variable, it should return an error too
	_, err := db.Exec(createEventsTable)
	if err != nil {
		panic("Could not create events table " + err.Error())
	}
}
