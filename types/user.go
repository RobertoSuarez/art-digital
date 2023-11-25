package types

import "time"

type APIUser struct {
	ID        uint      `json:"_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	Status    string    `json:"status"`
	Type      string    `json:"type"`
	CountryID uint      `json:"country_id"`
	Birthday  time.Time `json:"birthday"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
