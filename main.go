package main

import (
	"fmt"
	"main/internal/bill"
	moneyregister "main/internal/money_register"
	paymentprocessor "main/internal/payment_processor"
	"time"
)

func main() {
	proc := paymentprocessor.CardPaymentProcessor{}
	fmt.Println("Processor created")
	card := paymentprocessor.Card{
		CardNumber: "123455652342342",
		Balance:    10000,
	}
	fmt.Println("Card created")

	b := bill.Bill{
		ID:          0,
		Amount:      100,
		Description: "lorem ipsum dolor sit amet, consectetur adipis",
		DueDate:     time.Now().Add(time.Hour),
		Paid:        false,
		PaymentType: &proc,
	}

	reg := moneyregister.MoneyRegister{}

	testPaymentAndRefund(&reg, &b, &card, &proc)
}

func testPaymentAndRefund(reg *moneyregister.MoneyRegister, b *bill.Bill, card paymentprocessor.PaymentMethod, proc *paymentprocessor.CardPaymentProcessor) {
	fmt.Println("Bill created")
	fmt.Println("--------------------------------")
	fmt.Println("Starting Processing payment")
	fmt.Println("--------------------------------")

	msg, err := reg.MakePayment(b, card)
	if err != nil {
		fmt.Printf("Payment failed: %v - %v\n", msg, err)
	} else {
		fmt.Printf("Payment processed successfully: %v\n", msg)
	}
	fmt.Println("--------------------------------")

	fmt.Println("Test repeated payment")
	msg, err = reg.MakePayment(b, card)
	if err != nil {
		fmt.Printf("Repeated payment failed: %v - %v\n", msg, err)
	} else {
		fmt.Printf("Repeated payment result: %v\n", msg)
	}

	fmt.Println("--------------------------------")
	fmt.Println("Test refund")

	msg, err = reg.Refund(b, card)

	if err != nil {
		fmt.Printf("Refund failed: %v - %v\n", msg, err)
	} else {
		fmt.Printf("Refund processed successfully: %v\n", msg)
	}

	fmt.Println("--------------------------------")
	fmt.Println("Test refund on unpaid bill")
	unpaidBill := bill.Bill{
		ID:          1,
		Amount:      100,
		Description: "lorem ipsum dolor sit amet, consectetur adipis",
		DueDate:     time.Now().Add(time.Hour),
		Paid:        false,
		PaymentType: proc,
	}
	fmt.Println("--------------------------------")
	msg, err = reg.Refund(&unpaidBill, card)

	if err != nil {
		fmt.Printf("Refund on unpaid bill failed: %v - %v\n", msg, err)
	} else {
		fmt.Printf("Refund on unpaid bill result: %v\n", msg)
	}

}
