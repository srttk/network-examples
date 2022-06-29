package utils

import "log"

func CheckError(e error, msg string) {

	if e != nil {
		log.Fatal("Error ", msg)
	}
}
