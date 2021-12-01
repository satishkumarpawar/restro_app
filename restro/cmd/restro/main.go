package main

import (
	"restro/cmd/config/db"
	"restro/httpio"
)

var DB = db.DConn()

func main() {
	DB.AutoMigrate(&httpio.User{})
	DB.AutoMigrate(&httpio.Restaurant{})
	DB.AutoMigrate(&httpio.Dish{})

	HandlerRouting()
}
