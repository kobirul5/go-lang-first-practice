package main

import "fmt"

type PaymentMethod interface {
	pay(amount float64)
}

type Bkash struct {
	apiKey string
}

func (bk Bkash) pay(amount float64) {
	// Implementation for Bkash payment processing
	fmt.Printf("paying %f tk with Bkash" , amount)
}

type PaymentService struct {
	method PaymentMethod
}

func NewPaymentService(method PaymentMethod) *PaymentService {
	return &PaymentService{method: method}
}

func (ps PaymentService) Checkout(){
	// ps.method = Bkash{apiKey: "your-api-key"}
	ps.method.pay(100.0)
}

func main() {
	bkash := Bkash{apiKey: "your-api-key"}
	ps := NewPaymentService(bkash)
	ps.Checkout()
}
