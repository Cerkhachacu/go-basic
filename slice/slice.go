package main

import "log"

func main() {
	var cars = []string{"BMW", "TOYOTA", "HONDA", "YAMAHA"}

	log.Println(cars[0])

	var name = "bambang pamungkas"
	log.Println(name[8:])
	log.Println(name[:7])
	log.Println(name[5:11])

	carsitem := cars[:2]

	log.Println(carsitem)
	log.Println(len(carsitem))
	log.Println(cap(carsitem))

	carsitem = append(carsitem, "DAIHATSU")

	log.Println(carsitem)
	log.Println(len(carsitem))
	log.Println(cap(carsitem))
	log.Println(cars)

	var new = []string{"omama"}
	var newL = []string{"olala", "balala"}

	newAll := copy(newL, new)
	log.Println(new)
	log.Println(newL)
	log.Println(newAll)
}
