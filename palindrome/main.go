package main

import "fmt"

func isPalindrome(arr []string) bool {
	f_pointer := 0
	s_pointer := len(arr) - 1

	for f_pointer < s_pointer {
		if arr[f_pointer] != arr[s_pointer] {
			return false
		}
		f_pointer++
		s_pointer--
	}
	return true
}

func main() {
	var goodbye string = "arrivederci"
	arr := make([]string, len(goodbye))

	for i, char := range goodbye {
		arr[i] = string(char)
	}

	fmt.Println(isPalindrome(arr))
}
