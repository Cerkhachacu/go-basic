package main

import (
	"fmt"
	"log"
)

// declare struct
type student struct {
	name   string
	kelas  string
	number int
}

func main() {

	// using normal struct
	var item student
	item.name = "bambang"
	item.kelas = "1"

	fmt.Println(item.name)
	fmt.Println(item.kelas)
	fmt.Println(item.number)

	// using struct with object
	data := student{
		name:   "saskia",
		kelas:  "31",
		number: 10,
	}

	log.Println(data)

	// pointer struct
	var c *student

	// add value pointer struct
	c = &student{
		name: "olala",
	}

	d := c
	log.Println(d)
	c.name = "new name"
	log.Println(d)
}
