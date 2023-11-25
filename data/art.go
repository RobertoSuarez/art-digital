package data

import (
	"time"

	"gorm.io/gorm"
)

type StatusArt string

const (
	Private StatusArt = "private"
	Public  StatusArt = "public"
)

type Art struct {
	gorm.Model
	Title           string
	Description     string
	Value           float64
	PublicationDate time.Time
	UserID          uint
	User            User
	Status          StatusArt
}
