package main

import (
	"log"
	"strings"
)

// func as a paramater
func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, p := range data {
		if filtered := callback(p); filtered {
			result = append(result, p)
		}
	}

	return result
}

func main() {
	var data = []string{"omama", "olala", "ethan"}
	var arrData = filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})

	log.Println(arrData)
}
