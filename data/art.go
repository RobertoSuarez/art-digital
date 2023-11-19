package data

import (
	"time"

	"gorm.io/gorm"
)

type StatusArt string

const (
	Public  StatusArt = "public"
	Private StatusArt = "private"
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
