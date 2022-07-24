package utils

import "log"

func ValidErrorPanic(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
