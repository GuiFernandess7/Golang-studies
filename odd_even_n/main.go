package main

import "fmt"

func main() {
	var arr [4]int = [4]int{1, 2, 3, 4}
	const n int = 100
	var odd_number []int

	for i := 0; i < n; i++ {
		if i % 2 == 0 {
			fmt.Println(i, "is a odd number")
			odd_number = append(odd_number, i)
		} else {
			fmt.Println(i, "is an even number")
		}
	}

	for index, value := range arr {
		fmt.Println(index, value)
	}

	fmt.Println(odd_number)

}