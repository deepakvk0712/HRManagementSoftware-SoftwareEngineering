package utils

import "os"

func SetEnvironmentVariables() {
	os.Setenv("dbPath", "../Database/HR_DB")
	os.Setenv("emailDomain", "@hrtool.com")
}
