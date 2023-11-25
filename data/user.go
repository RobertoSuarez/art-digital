package data

import (
	"time"

	"gorm.io/gorm"
)

type StatusUser string

const (
	Actived StatusUser = "actived"
	Blocked StatusUser = "blocked"
)

type TypeUser string

const (
	NormalUser TypeUser = "normal_user"
	Superuser  TypeUser = "superuser"
)

type User struct {
	gorm.Model
	Name      string
	Email     string
	Password  string
	Status    StatusUser
	TypeUser  TypeUser
	CountryID uint
	Country   Country
	Birthday  time.Time
}
