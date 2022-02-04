package utils

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

var Db *gorm.DB

func ConnectToDb() {
	var err error

	Db, err = gorm.Open(sqlite.Open(os.Getenv("dbPath")), &gorm.Config{})

	if err != nil {
		fmt.Println(err)

		os.Exit(3)
	}
}
