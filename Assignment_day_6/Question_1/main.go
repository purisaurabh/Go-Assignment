package main

import "fmt"

func printValue(s string, alice chan string, bob chan string, flag chan int) {

	str := ""

	for i := 0; i < len(s); i++ {
		ch := fmt.Sprintf("%c", s[i])

		if ch == "$" {
			alice <- str
			str = ""

		} else if ch == "#" {
			bob <- str
			str = ""

		} else if ch == "^" {
			flag <- 1
			break
		} else {
			str += ch
		}

	}

}

func main() {
	s := "helloBob$helloalice#howareyou?#Iamgood.howareyou?$^"

	alice := make(chan string)
	bob := make(chan string)
	flag := make(chan int)

	go printValue(s, alice, bob, flag)

	for {
		select {
		case val := <-alice:
			fmt.Printf("alice:%s,", val)

		case val := <-bob:
			fmt.Printf("bob:%s,", val)

		case <-flag:
			return
		default:
			break
		}

	}

}
