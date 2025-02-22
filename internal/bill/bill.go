package bill

import (
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

func (b *Bill) Pay(method paymentprocessor.PaymentMethod) error {
	if b.Paid {
		return fmt.Errorf("the bill is already paid")
	}

	if method.GetBalance() < b.Amount {
		return fmt.Errorf("insufficient balance: have %.2f, need %.2f", method.GetBalance(), b.Amount)
	}

	b.PaymentType.Pay(method)
	b.Paid = true
	b.PaymentDate = time.Now()
	return nil
}

func (b *Bill) Refund(method paymentprocessor.PaymentMethod) error {
	if !b.Paid {
		return fmt.Errorf("cannot refund bill ID %d: bill is not paid", b.ID)

	}
	fmt.Println("Refunding bill ID:", b.ID)
	b.PaymentType.Refund(method)
	b.Paid = false
	b.PaymentDate = time.Time{}
	return nil
}
