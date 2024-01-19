package main

import "fmt"

func findTheIndexSecondApproach() {
	m := map[int]string{
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday",
		7: "Sunday",
	}

	var index int
	fmt.Println("Enter the index")
	fmt.Scanln(&index)

	if index < 1 || index > 7 {
		fmt.Println("Not a day")
	} else {
		fmt.Println("The day for the given index is : ", m[index])
	}

}

func findTheIndexFirstApproach() {

	m := map[int]string{
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday",
		7: "Sunday",
	}

	var index int
	fmt.Println("Enter the index")
	fmt.Scanln(&index)

	// One Aproach

	value, flag := m[index]

	if flag == true {
		fmt.Println("The day for the following index is :", value)
	} else {
		fmt.Println("Not a day")
	}

}

func main() {
	// findTheIndexFirstApproach()
	findTheIndexSecondApproach()
}
