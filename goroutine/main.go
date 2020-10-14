package main

import (
	"fmt"
	"log"
	"time"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		log.Println((i + 1), message)
		time.Sleep(2)
	}
}

func main() {
	go print(5, "hello")
	print(5, "apa kabs")

	var input string
	fmt.Scanln(&input)
}
