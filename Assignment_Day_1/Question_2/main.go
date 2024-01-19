/*
	2. Write a program to calculate the area of the circle, First line has a value of the radius of the circle
	constraint
	1. Use const PI from the package math package
	2. Use the Pow function from the math package
	3. Round area float value to 2 decimal places

*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64 = 9

	area := (math.Pi * math.Pow(radius, 2))
	result := fmt.Sprintf("%.2f", area)

	fmt.Println("Area of circle is ", result)
}
