package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func accessSlice(index int, nums []int) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Internal Error : %v\n", r)
		}
	}()

	// handling the negative input
	if index < 0 {
		panic("Index value can't ne negative")
	} else if index >= len(nums) {
		panic("Index Out Of Range")
	} else {
		fmt.Printf("Index : %d , value : %d ", index, nums[index])
	}
}

func main() {
	var index int
	nums := []int{1, 2, 3}
	fmt.Print("Enter the index : ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	index, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil {
		fmt.Println("Enter Correct Input Format")
		return
	}

	accessSlice(index, nums)
	fmt.Println("Testing panic and recover")
}
