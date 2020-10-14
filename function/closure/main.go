package main

import "log"

func main() {
	// closure func
	var closure = func(isError bool, message string) string {
		var status string = "Success Message"
		if !isError {
			status = "Error Message"
		}

		status = status + ": " + message
		return status
	}
	// (true, "oke nihh")

	// log.Println(closure)

	log.Println(closure(false, "is not oke"))
}
