package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func accessSlice(index int, nums []int) (int, error) {

	if index < 0 || index > len(nums)-1 {
		return 0, errors.New("length of the slice should be more than index")
	}

	return nums[index], nil

}

func main() {

	// for getting the length of the slice
	var n int
	fmt.Print("Enter length of slice : ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	n, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil {
		fmt.Println("Enter Correct Input Format")
		return
	}

	// for geting the elements of slice

	nums := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Print("Enter the elemets of index: ", i, " : ")
		input, _ = reader.ReadString('\n')
		value, err := strconv.Atoi(strings.TrimSpace(input))

		if err != nil {
			fmt.Println("Enter Correct Input Format")
			return
		}

		nums[i] = value
	}

	// for getting the index
	var index int
	fmt.Print("Enter the index : ")
	input, _ = reader.ReadString('\n')
	index, err = strconv.Atoi(strings.TrimSpace(input))

	if err != nil {
		fmt.Println("Enter Correct Input Format")
		return
	}

	result, err := accessSlice(index, nums)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Index : %d Value : %d\n", index, result)

}
