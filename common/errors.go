package common

import "log"

// CheckError is used for log error
// Return true if error not nil
func CheckError(err error) bool {
	if err != nil {
		log.Println("error:", err)
		return true
	}
	return false
}

// CheckErrorFatal is used for log error and exit
func CheckErrorFatal(err error) {
	if err != nil {
		log.Fatal("fatal error: ", err)
	}
}

// CheckErrorMessage is used for log error with custom message
// Return true if error not nil
func CheckErrorMessage(message string, err error) bool {
	if err != nil {
		log.Println(message, err)
		return true
	}
	return false
}
