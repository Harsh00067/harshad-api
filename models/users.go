package models

import "time"

type User struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	DOJ        string    `json:"doj"`
	Created_AT time.Time `json:"created_at"`
	Updated_AT time.Time `json:"updated_at"`
}
