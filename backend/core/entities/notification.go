package entities

import "time"

type Notification struct {
	Id string `json:"id"`
	Type string `json:"type"`
	Message string `json:"message"`
	At time.Time `json:"at"`
	From string `json:"from"`
	To string `json:"to"`
}
