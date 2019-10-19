package models

import "time"

/**
 * Entity Game
 */
type Game struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Platform    string    `json:"platform"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreateAt    time.Time `json:"createat"`
	UpdateAt    time.Time `json:"updateat"`
}
