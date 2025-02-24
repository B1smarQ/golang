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
	bankProc := paymentprocessor.BankTransferPaymentProcessor{}
	cashProc := paymentprocessor.CashPaymentProcessor{}
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

	bankAccount := paymentprocessor.BankAccount{
		Balance: 1000,
	}

	reg := moneyregister.MoneyRegister{}

	testPayment(&reg, &b, &card, &proc)
	b = bill.Bill{
		ID:          1,
		Amount:      100,
		Description: "lorem ipsum dolor sit amet, consectetur adipis",
		DueDate:     time.Now().Add(time.Hour),
		Paid:        false,
		PaymentType: &bankProc,
	}
	fmt.Println("--------------------------------")
	fmt.Println("--------------------------------")
	testPayment(&reg, &b, &bankAccount, &bankProc)
	fmt.Println("--------------------------------")
	fmt.Println("--------------------------------")
	cash := paymentprocessor.Cash{
		Balance: 1000,
	}
	b = bill.Bill{
		ID:          2,
		Amount:      100,
		Description: "lorem ipsum dolor sit amet, consectetur adipis",
		DueDate:     time.Now().Add(time.Hour),
		Paid:        false,
		PaymentType: &cashProc,
	}
	testPayment(&reg, &b, &cash, &cashProc)
}

func testPayment(reg *moneyregister.MoneyRegister, b *bill.Bill, card paymentprocessor.PaymentMethod, proc paymentprocessor.PaymentProcessor) {
	fmt.Println("Bill created")
	fmt.Println("--------------------------------")
	fmt.Println("Starting Processing payment")
	fmt.Println("--------------------------------")

	msg, err := reg.MakePayment(b, card)
	if err != nil {
		fmt.Println("Payment failed:", msg, "-", err)
	} else {
		fmt.Println("Payment processed successfully:", msg)
	}
	fmt.Println("--------------------------------")

	fmt.Println("Test repeated payment")
	msg, err = reg.MakePayment(b, card)
	if err != nil {
		fmt.Println("Repeated payment failed:", msg, "-", err)
	} else {
		fmt.Println("Repeated payment result:", msg)
	}

	fmt.Println("--------------------------------")
	fmt.Println("Test refund")

	msg, err = reg.Refund(b, card)

	if err != nil {
		fmt.Println("Refund failed:", msg, "-", err)
	} else {
		fmt.Println("Refund processed successfully:", msg)
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
		fmt.Println("Refund on unpaid bill failed:", msg, "-", err)
	} else {
		fmt.Println("Refund on unpaid bill result:", msg)
	}

	fmt.Println("--------------------------------")
}
