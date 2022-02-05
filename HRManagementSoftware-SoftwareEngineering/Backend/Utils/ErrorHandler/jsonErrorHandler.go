package errorHandler

import (
	"fmt"
	"os"
)

func JsonErrorHandler(err error, msg string) {
	if err != nil {
		fmt.Println(err)
		fmt.Println(msg)

		os.Exit(3)
	}
}
