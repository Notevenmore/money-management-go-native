package models

import "time"

type Debt struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Nominal  int       `json:"nominal"`
	Deadline time.Time `json:"deadline"`
	IsFinish bool      `json:"is_finish"`
}
