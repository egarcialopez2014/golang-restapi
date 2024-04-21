package models

import (
	"time"

	"example.com/apirest/db"
)

type Event struct { // this defines the structure of an event
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (e Event) Save() error {
	// later: add it to a database
	insertQuery := `
	INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(insertQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	// Exec for queries that modify things
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	e.ID = id
	events = append(events, e)
	return err
}

func GetAllEvents() ([]Event, error) {
	fetchQuery := "SELECT * FROM events"

	rows, err := db.DB.Query(fetchQuery)
	// Query for queries that retrieve things
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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
