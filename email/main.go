package main

import "fmt"

func (e email) cost() float64 {
	if !e.isSubscribed {
		return 0.05 * float64(len(e.body))
	}
	return 0.01 * float64(len(e.body))
}

func (e email) format() {
	fmt.Println(e.body)
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}
