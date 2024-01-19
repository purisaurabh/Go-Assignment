package main

import (
	"fmt"
)

func main() {

	var roman string
	fmt.Print("Enter the Roman Value : ")
	fmt.Scan(&roman)

	if roman == "" {
		fmt.Println("Please enter correct roman value")
		return
	}

	decimal := 0
	var currVal int
	prevVal := -1

	for i := len(roman) - 1; i >= 0; i-- {

		if roman[i] == 'I' {
			currVal = 1
		} else if roman[i] == 'V' {
			currVal = 5
		} else if roman[i] == 'X' {
			currVal = 10
		} else if roman[i] == 'L' {
			currVal = 50
		} else if roman[i] == 'C' {
			currVal = 100
		} else if roman[i] == 'D' {
			currVal = 500
		} else if roman[i] == 'M' {
			currVal = 1000
		} else {
			decimal = -1
			break
		}

		if prevVal > currVal {
			decimal -= currVal
		} else {
			decimal += currVal
		}

		prevVal = currVal
	}

	if decimal != -1 {
		fmt.Println("decimal value is", decimal)
	} else {
		fmt.Println("Please enter correct roman value")
	}
}
