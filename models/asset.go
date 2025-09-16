package models

type Assets struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	IsReusable bool   `json:"is_reusable"`
}
