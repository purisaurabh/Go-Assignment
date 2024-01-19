package main

import "fmt"

func returnTheSlice() {
	s := []string{"qwe", "wer", "ert", "rty", "tyu", "yui", "uio", "iop"}

	var index1, index2 int
	fmt.Println("Enter the two indices")
	fmt.Scan(&index1, &index2)

	if index1 < 0 || index2 < 0 || index1 >= len(s) || index2 >= len(s) {
		fmt.Println("Incorrect Indexes")
	} else {
		fmt.Println(s[:index1+1])
		fmt.Println(s[index1 : index2+1])
		fmt.Println(s[index2:])
	}

}

func main() {
	returnTheSlice()
}
