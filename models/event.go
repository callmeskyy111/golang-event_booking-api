package models

import "time"

// shape of the evt.
type Event struct {
	ID          int
	Name        string	`binding:"required"`
	Description string	`binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId 		int
}

var events = []Event{}

func(e Event) Save(){
	// later: Add to DB
	events = append(events, e)
}

func GetAllEvents() []Event{
return events
}