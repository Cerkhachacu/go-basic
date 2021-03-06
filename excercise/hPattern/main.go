package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// request input number
	reader1 := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter odd number bigger than 1 : ")
	scanner1 := reader1.Scan() // scan input
	reader2 := bufio.NewScanner(os.Stdin)
	fmt.Print("Chose method 1, 2 or 3 : ")
	scanner2 := reader2.Scan() // scan input
	if scanner1 {              //if input exist
		number, err := strconv.ParseInt(reader1.Text(), 10, 32) // conver input into integer
		if err == nil {
			if scanner2 {
				methodChosen, err := strconv.ParseInt(reader2.Text(), 10, 32)
				if err == nil {
					switch methodChosen {
					case 1:
						fmt.Print(patternH1(number)) // calling function to create the pattern
					case 2:
						patternH2(number)
					case 3:
						patternH3(number)
					default:
						fmt.Println("Please chose between method 1 or 2.")
					}
				} else {
					fmt.Println("Invalid input, only number is accepted")
				}
			} // success convert to integer
		} else { // if the input is not a number
			fmt.Println("Invalid input, only number is accepted")
		}
	}
}

func patternH1(number int64) string { // success convert to integer
	output := ""
	var numberCu int64 = number*number + 1
	var i int64 = 1
	mid := math.Ceil(float64(number) / 2)
	var r float64 = 1
	firstCol := true
	if number%2 != 0 && number > 1 { // cek if input is odd number
		for i < numberCu { // as long as
			if i%number != 0 { // if i modulus number equal 0
				if firstCol {
					output += "* "
					firstCol = false
				} else {
					if r == mid {
						output += "* "
						firstCol = false
					} else {
						output += "= "
					}
				}
			} else {
				r++
				output += "*\n"
				firstCol = true
			}
			i++
		}
	} else { // if the input is not an odd number print error
		output = "Please input an odd number that is bigger than one"
	}
	return output
}

func patternH2(number int64) { // success convert to integer
	if number%2 != 0 && number > 1 { // cek if input is odd number
		var i float64 = 1
		mid := math.Ceil(float64(number) / 2)
		for i <= float64(number) {
			if mid != i { // creating "* = = = *" part
				starSign := "* *\n"
				equalSign := strings.Repeat(" =", int(number-2))
				output := starSign[:1] + equalSign + starSign[1:]
				fmt.Print(output)
			} else { // creating "* * * * *"
				starSign := "* *\n"
				equalSign := strings.Repeat(" *", int(number-2))
				output := starSign[:1] + equalSign + starSign[1:]
				fmt.Print(output)
			}
			i++
		}
	} else { // if the input is not an odd number print error
		fmt.Println("Please input an odd number that is bigger than one")
	}
}

func patternH3(number int64) { // success convert to integer
	if number%2 != 0 && number > 1 { // cek if input is odd number
		var i float64 = 1
		var j float64 = 2
		mid := math.Ceil(float64(number) / 2)
		for i <= float64(number) {
			fmt.Print("* ")
			for j <= float64(number) {
				if i == mid {
					if j < float64(number) {
						fmt.Print("* ")
					} else {
						fmt.Print("*\n")
					}
				} else {
					if j != float64(number) {
						fmt.Print("= ")
					} else {
						fmt.Print("*\n")
					}
				}
				j++
			}
			j = 2
			i++
		}
	} else { // if the input is not an odd number print error
		fmt.Println("Please input an odd number that is bigger than one")
	}
}
