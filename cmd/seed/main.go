package main

import (
	"fmt"

	"github.com/RobertoSuarez/art-digital/data"
	"github.com/RobertoSuarez/art-digital/db"
)

func main() {
	fmt.Println("Poblando la base de datos")

	db.Init()
	db.DB.AutoMigrate(&data.User{})

	db.DB.Create(&data.User{
		Name:  "Carlos",
		Email: "carlos@gmail.com",
	})

	db.DB.Create(&data.User{
		Name:  "Daniel",
		Email: "daniel@gmail.com",
	})
}
