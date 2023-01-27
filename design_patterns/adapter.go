package main

import "fmt"

type IPayment interface {
	Pay()
}

type CashPayment struct{}

func (CashPayment) Pay() {
	fmt.Println("Payment using cash")
}

func ProcessPayment(p IPayment) {
	p.Pay()
}

type TransferPayment struct{}

func (TransferPayment) Pay(bankAccount string) {
	fmt.Println("Payment using transfer")
	fmt.Printf("Using the account %s", bankAccount)
}

type TransferPaymentAdapter struct {
	TransferPayment *TransferPayment
	BankAccount     string
}

func (tpa *TransferPaymentAdapter) Pay() {
	fmt.Println("Payment using Transfer and an adapter")
	tpa.TransferPayment.Pay(tpa.BankAccount)
}

func main() {
	cash := &CashPayment{}
	ProcessPayment(cash)

	transferAdapter := &TransferPaymentAdapter{
		TransferPayment: &TransferPayment{},
		BankAccount:     "1111222",
	}
	ProcessPayment(transferAdapter)
}
