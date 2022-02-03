package db

import (
	"database/sql"
	"fmt"
	"os"

	utils "github.com/dstejas19/HRManagementSoftware-SoftwareEngineering/Backend/Utils"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func ConnectToDB() {
	var err error
	Db, err = sql.Open("sqlite3", utils.Paths.DbPath)

	if err != nil {
		fmt.Println(err)

		os.Exit(3)
	}
}
