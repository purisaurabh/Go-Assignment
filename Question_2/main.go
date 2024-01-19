package main

import (
	"bufio"
	"fmt"
	"os"
)

func wordCount() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the string : ")
	scanner.Scan()
	str := scanner.Text()

	m := make(map[string]int, 0)
	temp := make([]string, 0)
	s := ""

	for i := 0; i < len(str); i++ {

		ch := fmt.Sprintf("%c", str[i])

		if ch != " " {
			s = s + ch
		} else {
			temp = append(temp, s)
			m[s]++
			s = ""
		}
	}

	if s != "" {
		m[s]++
		temp = append(temp, s)
	}

	maxFreq := 0

	for _, val := range m {
		if val > maxFreq {
			maxFreq = val
		}
	}

	ans := make([]string, 0)

	for i := 0; i < len(temp); i++ {
		if m[temp[i]] == maxFreq {
			ans = append(ans, temp[i])
			m[temp[i]] = -1
		}
	}

	fmt.Println(ans)

}

func main() {
	wordCount()
}
