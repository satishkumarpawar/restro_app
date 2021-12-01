package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

var urlIDSN string = "skp:123@tcp(localhost:3306)/restrodb?parseTime=true"
var err error

func DConn() *gorm.DB {
	Database, err = gorm.Open(mysql.Open(urlIDSN), &gorm.Config{})
	if err != nil {
		fmt.Print(err.Error())
		panic("connection failed :(")
	}
	return Database

}
