package main

import (
	"fmt"
	"time"
)

var payments []int = []int{}

// Methods
type PaymentMethod interface {
	pay(amount float32) string
}

// Objects
type CreditCard struct {
	number string
	cvv int
	expiry time.Time
}

type Paypal struct {
	email string
}

// Implementações
func (card CreditCard) pay(value float32) bool {
	payments = append(payments, int(value))
	fmt.Println("Payment is successfull!")
	return true
}

func (card Paypal) pay(value float32) bool {
	payments = append(payments, int(value))
	fmt.Println("Payment is successfull!")
	return true
}

func main(){
	var card CreditCard = CreditCard{
		number: "3252.9374.9261.1723",
		cvv: 211,
		expiry: time.Date(2031, 4, 20, 0, 0, 0, 0, time.UTC),
	}
	card.pay(200.0)
}