package moneyregister

import (
	PaymentBill "main/internal/bill"
	paymentprocessor "main/internal/payment_processor"
)

type MoneyRegister struct {
}

func (r *MoneyRegister) MakePayment(bill *PaymentBill.Bill, PaymentMethod paymentprocessor.PaymentMethod) {
	bill.Pay(PaymentMethod)
}

func (r *MoneyRegister) Refund(bill *PaymentBill.Bill, PaymentMethod paymentprocessor.PaymentMethod) {
	bill.Refund(PaymentMethod)
}
