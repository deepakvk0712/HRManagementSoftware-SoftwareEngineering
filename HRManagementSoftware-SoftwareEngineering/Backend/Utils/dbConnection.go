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
	fmt.Println(os.Getenv("DB_PATH"))
	Db, err = gorm.Open(sqlite.Open(os.Getenv("DB_PATH")), &gorm.Config{})

	if err != nil {
		fmt.Println(err)

		os.Exit(3)
	}
}
