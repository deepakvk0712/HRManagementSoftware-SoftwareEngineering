package utils

import "github.com/joho/godotenv"

func Init() {
	godotenv.Load("./secrets.env")
	ConnectToDb()

}
