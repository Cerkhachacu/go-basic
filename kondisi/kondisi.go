package main

import (
	"fmt"
)

func main() {
	// if esle
	item := 0
	if item == 0 {
		fmt.Println("nilai item: ", item)
	} else if item > 5 {
		fmt.Println("item lebih besar dari pada 5", item)
	} else {
		fmt.Println("item diantara 0 - 5")
	}

	// temporary variabel
	if allItem := item + 10; allItem > 20 {
		fmt.Println("nilai semua item lebih besar dari 20")
	} else {
		fmt.Println("nilai item kurang dari 20")
	}

	switch item {
	case 0:
		fmt.Println("olala")
	case 1:
		fmt.Println("omama")
	default:
		fmt.Println("horay")
	}

	switch item {
	case 0:
		fmt.Println("kosong")
	case 1, 2, 3, 4, 5:
		fmt.Println("good")
	default:
		fmt.Println("nice")
	}

	switch item {
	case 3:
		fmt.Println("balala")
	case 1, 2:
		fmt.Println("kalasa")
	default:
		{
			fmt.Println("omama")
			fmt.Println("babba")
		}
	}
}
