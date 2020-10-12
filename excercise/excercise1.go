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
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter odd number bigger than 1 : ")
	scanner := reader.Scan() // scan input
	if scanner {             //if input exist
		number, err := strconv.ParseInt(reader.Text(), 10, 32) // conver input into integer
		if err == nil {                                        // success convert to integer
			patternH1(number) // calling function to create the pattern
		} else { // if the input is not a number
			fmt.Println("Invalid input, only number is accepted")
		}
	}
}

func patternH1(number int64) { // success convert to integer
	if number%2 != 0 && number > 1 { // cek if input is odd number
		var numberCu int64 = number*number + 1
		var i int64 = 1
		mid := math.Ceil(float64(number) / 2)
		var r float64 = 1
		firstCol := true
		for i < numberCu { // as long as
			if i%number != 0 { // if i modulus number equal 0
				if firstCol {
					fmt.Print("* ")
					firstCol = false
				} else {
					if r == mid {
						fmt.Print("* ")
						firstCol = false
					} else {
						fmt.Print("= ")
					}
				}
			} else {
				r++
				fmt.Print("*\n")
				firstCol = true
			}
			i++
		}
	} else { // if the input is not an odd number print error
		fmt.Println("Please input an odd number that is bigger than one")
	}
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
