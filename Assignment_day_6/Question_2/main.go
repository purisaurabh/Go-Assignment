package main

import (
	"fmt"
	"runtime"
)

func reverseString(str string, revStr chan string) {

	val := <-revStr

	fmt.Println("The given string is : ", str)
	fmt.Println("The Reverse string is : ", val)
	fmt.Println("The number of goroutine are : ", runtime.NumGoroutine())
}

func main() {
	str := "JAMES"

	revStr := make(chan string, 0)

	go reverseString(str, revStr)

	dumpStr := ""

	for i := 0; i < len(str); i++ {
		ch := fmt.Sprintf("%c", str[i])
		dumpStr = ch + dumpStr
	}

	revStr <- dumpStr

}
