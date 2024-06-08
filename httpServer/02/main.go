package main

import (
	"fmt"
	"time"
)

func counter(t string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%s: %d\n", t, i)
		time.Sleep(time.Millisecond * 500)
	}
}

func workerCalculate(workerID string, msg chan int){
	for res := range msg {
		fmt.Printf("Worker %s received %d\n", workerID, res)
		//time.Sleep(time.Second)
	}
}

// T1
func main(){
	msg := make(chan int)

	go workerCalculate("User 1", msg)
	go workerCalculate("User 2", msg)

	for i := 0; i < 10; i++ {
		msg <- i
	}

	/* queue := make(chan int)

	// T2
	go func() {
		i := 0
		for {
			queue <- i
			i++
			time.Sleep(time.Second)
		}
	}()

	// T3
	for x := range queue {
		fmt.Println(x) // Esvazia o canal sempre ao rodar linha 22
	} */
}

