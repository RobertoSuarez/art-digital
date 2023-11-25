package types

import (
	"errors"
	"time"
)

type UserAPI struct {
	ID        uint      `json:"_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	Status    string    `json:"status"`
	Type      string    `json:"type"`
	CountryID uint      `json:"country_id"`
	Birthday  time.Time `json:"birthday"`
}

func (user *UserAPI) Validate() error {

	if !(len(user.Password) > 6) {
		return errors.New("la contraseña no es mayor a 6 caracteres")
	}

	return nil
}
