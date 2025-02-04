package models

import "time"

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	Date        time.Time `binding:"required"`
	UserID      int
}

var Events = []Event{}

// func (e Event) CreateEvent() {
// 	events := append(Events, e)
// }

func  GetAllEvent() []Event {
	return Events
}

// func (e Event) GetEvents(id int) Event {

// }
