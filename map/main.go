package main

import "log"

func main() {
	var month = map[string]int{"januari": 30, "febuari": 28, "maret": 30}
	log.Println(month["januari"])

	for index, val := range month {
		log.Println(index+" : ", val)
	}
	delete(month, "januari")
	var value, isExist = month["januari"]
	if isExist {
		log.Println(value)
	}
}
