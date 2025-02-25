package moneyregister

import (
	PaymentBill "main/internal/bill"
	paymentprocessor "main/internal/payment_processor"
)

type MoneyRegister struct {
}

func (r *MoneyRegister) MakePayment(bill *PaymentBill.Bill, PaymentMethod paymentprocessor.PaymentMethod, paymentProcessor paymentprocessor.PaymentProcessor) (string, error) {
	bill.PaymentType = paymentProcessor
	message, err := bill.Pay(PaymentMethod)
	return message, err
}

func (r *MoneyRegister) Refund(bill *PaymentBill.Bill, PaymentMethod paymentprocessor.PaymentMethod, paymentProcessor paymentprocessor.PaymentProcessor) (string, error) {
	bill.PaymentType = paymentProcessor
	message, err := bill.Refund(PaymentMethod)
	return message, err
}
