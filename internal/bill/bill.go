package bill

import (
	"errors"
	"fmt"
	paymentprocessor "main/internal/payment_processor"
	"time"
)

type Bill struct {
	ID          int
	Amount      float64
	Description string
	DueDate     time.Time
	Paid        bool
	PaymentType paymentprocessor.PaymentProcessor
	PaymentDate time.Time
}

func (b *Bill) Pay(method paymentprocessor.PaymentMethod) (string, error) {
	if b.Paid {
		return "Error", fmt.Errorf("the bill is already paid")
	}

	if method.GetBalance() < b.Amount {
		return "Error", fmt.Errorf("insufficient balance: have %.2f, need %.2f", method.GetBalance(), b.Amount)
	}

	b.PaymentType.Pay(method)
	b.Paid = true
	b.PaymentDate = time.Now()
	return "Success", nil
}

func (b *Bill) Refund(method paymentprocessor.PaymentMethod) (string, error) {
	if !b.Paid {

		return "Processing error", errors.New("the bill is not paid")

	}
	fmt.Println("Refunding bill ID:", b.ID)
	b.PaymentType.Refund(method)
	b.Paid = false
	b.PaymentDate = time.Time{}
	return "", nil
}
