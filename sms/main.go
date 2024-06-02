package main

import "fmt"

// Interfaces | Methods
type expense interface {
	cost() float64
}

// Structs | Objects
type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

// Implementaions | Getters
func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

// Helper Function
func getExpenseReport(e expense)(string, float64){
	em, ok := e.(email)
	if ok {
		return em.toAddress, em.cost()
	}
	sms, ok := e.(sms)
	if ok {
		return sms.toPhoneNumber, sms.cost()
	}
	return "", 0.0
}

func main(){
	email := email{
		isSubscribed: true,
		body: "Active",
		toAddress: "Alameda Street",
	}
	fmt.Println(getExpenseReport(email))
}