/*
	1. Write a program to calculate the simple interest
	First-line has the comma-separated values of the Principal, rate and time (in years) respective
	*constraints: Round simple interest float value to 2 decimal places
*/

package main

import "fmt"

func main() {

	principal, interestRate, timeInYears := 30000.0, 8.0, 5.4

	simpleInterest := (principal * interestRate * timeInYears) / 100

	result := fmt.Sprintf("%.2f", simpleInterest)
	fmt.Println("Simple Interest is : ", result)
}
