package db

import (
	"database/sql"
	"fmt"

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

func UpdateUserPassword(email, newPassword string) error {
	// Prepare the query to update the password
	query := "UPDATE users SET password = ? WHERE email = ?"
	_, err := DB.Exec(query, newPassword, email)
	if err != nil {
		return fmt.Errorf("failed to update password: %v", err)
	}
	fmt.Println("Password updated successfully")
	return nil
}

func GetUserPassword(email string) (string, error) {
	query := "SELECT password FROM users WHERE email = ?"
	var password string
	err := DB.QueryRow(query, email).Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("no user found with email %s", email)
		}
		return "", fmt.Errorf("failed to query password: %v", err)
	}
	return password, nil
}

//create tables for the event

func createTables(db *sql.DB) {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("Could not create users table")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	location TEXT NOT NULL,
	dateTime DATETIME NOT NULL,
	user_id INTEGER,
	FOREIGN KEY(user_id) REFERENCES users(id) 
	)`

	//To execute a query statement
	//_, err := DB.Exec(createEventsTable) // store event table in a variable, it should return an error too
	_, err = db.Exec(createEventsTable)
	if err != nil {
		panic("Could not create events table " + err.Error())
	}

	createRegistrationsTable := `
	 CREATE TABLE IF NOT EXISTS registrations (
	 id INTEGER PRIMARY KEY AUTOINCREMENT,
	 event_id INTEGER,
	 user_id INTEGER,
	 FOREIGN KEY(event_id) REFERENCES events(id),
	 FOREIGN KEY(user_id) REFERENCES users(id)
	 )
	 `
	_, err = db.Exec(createRegistrationsTable)
	if err != nil {
		panic("Could not create registrations table")
	}
}
