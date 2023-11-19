package data

import (
	"time"

	"gorm.io/gorm"
)

type StatusUser string

const (
	Blocked StatusUser = "blocked"
	Actived StatusUser = "actived"
)

type TypeUser string

const (
	NormalUser TypeUser = "normal_user"
	SuperUser  TypeUser = "superuser"
)

type User struct {
	gorm.Model
	Name        string
	Email       string
	Status      StatusUser
	Type        TypeUser
	CountryID   uint
	Country     Country
	DateOfBirth time.Time
}
