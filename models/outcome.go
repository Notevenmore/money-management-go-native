package models

import "time"

type Outcome struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Nominal  int       `json:"nominal"`
	Date     time.Time `json:"date"`
	Category string    `json:"category"`
}
