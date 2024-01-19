/*
	1. The given program accepts an integer value between 1 to 4 from the user.
	There is an option associated with each value, which is basically objects of different data types with some associated value.

	Write a method named "AcceptAnything" which takes any type of data as input.

	Based on the option chosen by the user, we will send different types of objects to this function and it should print the following based on its respective data type of value sent to the function "AcceptAnything".

	integer :
	"This is a value of type Integer, <value>"

	string :
	"This is a value of type String, <value>"

	boolean :
	"This is a value of type Boolean, <value>"

	Custom data type Hello :
	"This is a value of type Hello, <value>"
*/

package main

import "fmt"

type dataType interface {
}

type fourthCase struct {
	number int
	value  string
}

func AcceptAnything(d dataType) {
	fmt.Printf("The type for the given input is : %T and the respective value is %v\n", d, d)
}

func main() {
	fmt.Println("Question_1")
	var choice int
	fmt.Println(" 1.Integer \n 2.String \n 3.Boolean \n 4.Custom Data Type \n Enter a choice :")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		integervalue := 10
		AcceptAnything(integervalue)

	case 2:
		str := "string"
		AcceptAnything(str)

	case 3:
		flag := true
		AcceptAnything(flag)
	case 4:
		st := fourthCase{10, "ten"}
		AcceptAnything(st)
	default:
		fmt.Println("Invalid Case")
	}

}
