package models

import (
	"fmt"
	"time"

	"example.com/events-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

type Registration struct {
	ID        int64     `json:"id"`
	EventID   int64     `json:"event_id"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

// Save function that will save the events to a database
var events = []Event{}

func (e *Event) Save() error {
	// later add it to a database
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id) 
	VALUES (?,?,?,?,?)` // to input data into the data base
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()                                                                //closing the statement (stmt) after we have used it.
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID) // inserting data into the database
	if err != nil {
		return err
	}
	id, err := result.LastInsertId() // to get the ID on the event that was inserted in the database
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query) //fetch data from the data base
	if err != nil {
		return nil, err
	}
	defer rows.Close() //close after you get some rows from data
	//loop through all the rows to step by step read them and pouplate and event slice with that row data
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID) //to populate all the different field of that struct
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (event Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err
}
func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.ID)
	return err
}

func (e Event) Register(userId int64) error {
	query := "INSERT INTO registrations(event_id, user_id, created_at) VALUES (?, ?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID, userId, time.Now())
	return err
}

func (e Event) CancelRegistration(userId int64) error {
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID, userId)
	return err
}

func GetAllRegistrations() ([]Registration, error) {
	query := "SELECT * FROM registrations"
	rows, err := db.DB.Query(query) //fetch it from the database
	if err != nil {
		fmt.Printf("DEBUG: Query error: %v\n", err)
		return nil, fmt.Errorf("query error: %v", err)
	}
	defer rows.Close()
	// Debugging: Print column names returned by the query
	columns, err := rows.Columns()
	if err != nil {
		fmt.Printf("DEBUG: Columns error: %v\n", err)
		return nil, fmt.Errorf("columns error: %v", err)
	}
	fmt.Printf("DEBUG: Columns returned by query: %v\n", columns)
	var registrations []Registration
	for rows.Next() {
		var registration Registration
		err := rows.Scan(&registration.ID, &registration.EventID, &registration.UserID, &registration.CreatedAt)
		if err != nil {
			fmt.Printf("DEBUG: Row scan error: %v\n", err)
			return nil, fmt.Errorf("row scan error: %v", err)
		}
		registrations = append(registrations, registration)
	}

	// Check for errors during iteration
	if err = rows.Err(); err != nil {
		fmt.Printf("DEBUG: Row iteration error: %v\n", err)
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return registrations, nil
}
