package common

import "log"

// CheckError is used for log error
func CheckError(err error) {
	if err != nil {
		log.Println("error:", err)
	}
}

// CheckErrorFatal is used for log error and exit
func CheckErrorFatal(err error) {
	if err != nil {
		log.Fatal("fatal error: ", err)
	}
}

// CheckErrorMessage is used for log error with custom message
func CheckErrorMessage(message string, err error) {
	if err != nil {
		log.Println(message, err)
	}
}
