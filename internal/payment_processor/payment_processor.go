package paymentprocessor

import "fmt"

type PaymentProcessor interface {
	Pay(PaymentMethod)
	Refund(PaymentMethod)
}

type CashPaymentProcessor struct {
}

func (cp *CashPaymentProcessor) Pay() {
	fmt.Println("Paying with cash")
}

func (cp *CashPaymentProcessor) Refund() {
	fmt.Println("Refunding cash")
}

type CardPaymentProcessor struct {
}

func (cp *CardPaymentProcessor) Pay(p PaymentMethod) {
	if !cp.CheckVaditidy() {
		fmt.Println("Invalid card")
		return // Return early to avoid unnecessary processing
	}
	fmt.Println("Paying with card")
}

func (cp *CardPaymentProcessor) Refund(p PaymentMethod) {
	fmt.Println("Refunding card")
}

func (cp *CardPaymentProcessor) CheckVaditidy() bool {
	fmt.Println("Checking card validity")
	fmt.Println("Verification successful")
	return true // Replace with actual card validation logic
}

type BankTransferPaymentProcessor struct {
}

func (cp *BankTransferPaymentProcessor) Pay() {
	fmt.Println("Paying with bank transfer")
}

func (cp *BankTransferPaymentProcessor) Refund() {
	fmt.Println("Refunding bank transfer")
}
