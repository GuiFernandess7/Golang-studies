package main

import "fmt"

func fibGen() func() int {
    f1, f2 := 0, 1
    return func() int {
        f1, f2 = f2, f1+f2
        return f1
    }
}

func main(){
	f := fibGen()
	for i := 0; i < 10; i++ {
		fmt.Print(f(), " ")
	}
}