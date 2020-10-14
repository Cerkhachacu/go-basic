package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type triangle interface {
	rataKanan() string
	rataKiri() string
}

type inputNumber struct {
	number int64
}

func (number inputNumber) rataKanan() string {
	var triangle string
	var i = 1
	for i <= int(number.number) {
		triangle += strings.Repeat("* ", int(i)) + "\n"
		i++
	}
	return triangle
}

func (number inputNumber) rataKiri() string {
	var triangle string
	var i = 0
	for i < int(number.number) {
		if i == 0 {
			triangle += strings.Repeat("* ", int(number.number)-i) + "\n"
		} else {
			triangle += strings.Repeat(" ", i+i) + strings.Repeat("* ", int(number.number)-i) + "\n"
		}
		i++
	}
	return triangle
}

func main() {
	reader1 := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter the length of the Isosceles triangle : ")
	scanner1 := reader1.Scan() // scan input
	reader2 := bufio.NewScanner(os.Stdin)
	fmt.Print("Chose rataKanan (1) or rataKiri(2) : ")
	scanner2 := reader2.Scan() // scan input
	if scanner1 {              //if input exist
		number, err := strconv.ParseInt(reader1.Text(), 10, 32) // conver input into integer
		if err == nil {
			if scanner2 {
				patternChose, err := strconv.ParseInt(reader2.Text(), 10, 32)
				if err == nil {
					var segitiga triangle

					segitiga = inputNumber{number}
					switch patternChose {
					case 1:
						fmt.Println(segitiga.rataKanan()) // calling function to create the pattern
					case 2:
						fmt.Println(segitiga.rataKiri())
					default:
						fmt.Println("Please chose between method 1 or 2.")
					}
				} else {
					fmt.Println("Invalid input, only number is accepted")
				}
			}
		} else { // if the input is not a number
			fmt.Println("Invalid input, only number is accepted")
		}
	}
}
