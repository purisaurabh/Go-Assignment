/*

	3. The given program takes an integer value as input from the user i.e 1 or 2.
	Option 1 represents Rectangle and Option 2 represents Square.

	Given the metrics of the shapes are hard-coded, complete the given program to achieve the following:

	1. Create an interface Quadrilateral which has the following method signatures
	* Area() int
	* Perimeter() int

	2. Implement the Quadrilateral interface for the given shapes i.e Rectangle and Square

	3. Define a "Print" function which accepts any shape that implements Quadrilateral interface and Prints Area and Perimeter of shape in the following manner:

	"Area :  <value>"
	"Perimeter :  <value>"

*/

package main

import "fmt"

type Quadrilateral interface {
	Area() int
	Perimeter() int
}

type getLengthWidth struct {
	length int
	width  int
}

type getSide struct {
	side int
}

func (gv getLengthWidth) Area() int {
	return gv.length * gv.width
}

func (gv getLengthWidth) Perimeter() int {
	addition := gv.length + gv.width
	return 2 * addition
}

func Print(q Quadrilateral) {
	fmt.Println("The area for the given rectangle is : ", q.Area())
	fmt.Println("The perimeter for the given rectangle is : ", q.Perimeter())
}

func Rectangle() {
	var l, w int
	fmt.Println("Enter the value of length :")
	fmt.Scan(&l)
	fmt.Println("Enter the value of width :")
	fmt.Scan(&w)

	// for rectangle
	var rec Quadrilateral
	rec = getLengthWidth{l, w}

	Print(rec)
}

func (gs getSide) Area() int {
	return gs.side * gs.side
}

func (gs getSide) Perimeter() int {
	return 4 * gs.side
}

func Square() {
	var s int
	fmt.Println("Enter the value of side of square :")
	fmt.Scan(&s)

	var sq Quadrilateral
	sq = getSide{s}

	Print(sq)
}

func main() {
	fmt.Println("Enter an Input : \n 1.Rectangle \n 2.Square ")
	var n int
	fmt.Scan(&n)

	switch n {
	case 1:
		Rectangle()
		break

	case 2:
		Square()
		break

	default:
		fmt.Println("Invalid case")
		break
	}
}
