package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
)

func reverseString(str string, revStr chan string, wg *sync.WaitGroup) {

	wg.Add(1)
	val := <-revStr

	fmt.Println("The given string is : ", str)
	fmt.Println("The Reverse string is : ", val)
	fmt.Println("The number of goroutine are : ", runtime.NumGoroutine())
	wg.Done()
}

func main() {

	var wg sync.WaitGroup
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string : ")
	input, _ := reader.ReadString('\n')
	str := strings.TrimSpace(input)

	revStr := make(chan string)
	go reverseString(str, revStr, &wg)

	dumpStr := ""

	for i := 0; i < len(str); i++ {
		ch := fmt.Sprintf("%c", str[i])
		dumpStr = ch + dumpStr
	}
	revStr <- dumpStr
	wg.Wait()
}
