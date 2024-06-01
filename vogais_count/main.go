package main

import "fmt"

func contains(slice []rune, item rune) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func main(){
	var goodbye string = "arrivederci"
	var count int = 0;
	var vogais []rune = []rune{'a', 'e', 'i', 'o', 'u'}

	for _, letter := range goodbye {
		if contains(vogais, letter) {
			count ++
		}
	}

	fmt.Println(count)

}