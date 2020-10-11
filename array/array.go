package main

import "log"

func main() {
	var cars = [4]string{"BMW", "TOYOTA", "DAIHATSU"}
	log.Println(cars)
	log.Println(len(cars))

	var numb = [...]int{1, 2, 3, 4, 5, 6}
	log.Println(numb)
	log.Println(len(numb))

	var multi = [][]int{[]int{1, 2, 4, 5}, []int{3, 4, 5, 6}}
	log.Println(multi)
	log.Println(len(multi))

}
