package entities

import "time"

type Message struct {
	ID      int       `json:"id"`
	From    string    `json:"from"`
	To      string    `json:"to"`
	Content string    `json:"content"`
	Image   string    `json:"image"`
	SendAt  time.Time `json:"date"`
}
