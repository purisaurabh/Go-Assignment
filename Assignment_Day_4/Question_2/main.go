/*
	2. The given program accepts 2 values from the user, length and breadth of a rectangle respectively.

	Complete the program by writing methods "Area" and "Perimeter" on Rectangle type to calculate the area and perimeter of a rectangle.

	Hint:
	Method Signatures for Area and Perimeter
	Area() int
	Perimeter() int

	Formulae:
	Area = length * width
	Perimeter = 2 * (length + width)

	Example Test Case 1:
	Input: 10 20
	Output:
	Area of Rectangle: 200
	Perimeter of Rectangle: 60

*/

package main

import "fmt"

type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) getArea() int {
	return int(r.length) * int(r.width)
}

func (r Rectangle) getPerimeter() int {
	addition := r.length + r.width
	return 2 * int(addition)
}

func main() {
	fmt.Println("Questin Second")

	var l, w float64
	fmt.Println("Enter the value of length :")
	fmt.Scanln(&l)

	fmt.Println("Enter the value of breadth :")
	fmt.Scanln(&w)

	res := Rectangle{l, w}

	fmt.Println("The value of length and breadth is : ", res.length, res.width)
	fmt.Println("Area of Rectangle : ", res.getArea())
	fmt.Println("The perimeter of rectangle : ", res.getPerimeter())

}
