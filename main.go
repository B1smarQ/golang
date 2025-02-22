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

	fmt.Println("Bill created")
	fmt.Println("--------------------------------")
	fmt.Println("Starting Processing payment")
	fmt.Println("--------------------------------")

	reg.MakePayment(&b, &card)
	fmt.Println("Payment processed")
	fmt.Println("--------------------------------")

	fmt.Println("Test repeated payment")
	tmp1, tmp2 := reg.MakePayment(&b, &card)
	fmt.Println(tmp1, tmp2)

	fmt.Println("--------------------------------")
	fmt.Println("Test refund")

	reg.Refund(&b, &card)
	fmt.Println("--------------------------------")
	fmt.Println("Test refund on unpaid bill")
	unpaidBill := bill.Bill{
		ID:          1,
		Amount:      100,
		Description: "lorem ipsum dolor sit amet, consectetur adipis",
		DueDate:     time.Now().Add(time.Hour),
		Paid:        false,
		PaymentType: &proc,
	}
	msg, err := reg.Refund(&unpaidBill, &card)
	if err != nil {
		fmt.Println(msg, err)
	}

}
