package models

import (
	"fmt"
	"time"

	"example.com/events-api/db"
)

type Registration struct {
	ID        int64     `json:"id"`
	EventID   int64     `json:"event_id"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
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

// get a single registration
func GetRegistrationByID(id int64) (*Registration, error) {
	query := "SELECT * FROM registrations WHERE id = ?"
	row := db.DB.QueryRow(query, id)
	var registration Registration
	err := row.Scan(&registration.ID, &registration.EventID, &registration.UserID, &registration.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &registration, nil

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
