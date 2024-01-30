package main

import (
	"fmt"
	"sync"
	"time"
)

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	var wg sync.WaitGroup
	// var m sync.Mutex

	n := 0

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// m.Lock()
		nIsEven := isEven(n)
		time.Sleep(5 * time.Millisecond)
		if nIsEven {
			fmt.Println(n, " is even")
			return
		}
		fmt.Println(n, "is odd")
		// m.Unnlock()

	}(&wg)
	wg.Add(1)
	wg.Wait()

	go func() {
		// m.Lock()
		n++
		// m.Unlock()
	}()

	// just waiting for the goroutines to finish before exiting
	time.Sleep(time.Second)

}
