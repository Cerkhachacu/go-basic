package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	// request input number
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter odd number bigger than 1 : ")
	scanner := reader.Scan() // scan input
	if scanner {             //if input exist
		number, err := strconv.ParseInt(reader.Text(), 10, 32) // conver input into integer
		if err == nil {                                        // success convert to integer
			if number%2 != 0 && number > 1 { // cek if input is odd number
				var numberCu int64 = number*number + 1
				var i int64 = 1
				mid := math.Round(float64(number) / 2)
				var r float64 = 1
				firstCol := true
				for i < numberCu { // creating the input
					if i%number != 0 {
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
				fmt.Println("Please input odd number that is bigger than one")
			}
		} else { // if the input is not a number
			fmt.Println("Invalid input, only number is accepted")
		}
	}
}
