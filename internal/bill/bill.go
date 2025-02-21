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

func (b *Bill) Pay(method paymentprocessor.PaymentMethod) {
	if !b.Paid {
		b.PaymentType.Pay(method)
		b.Paid = true
		b.PaymentDate = time.Now()
	} else {
		fmt.Println("The bill is already paid")
	}
}

func (b *Bill) Refund(method paymentprocessor.PaymentMethod) {
	if b.Paid {
		fmt.Println("Refunding bill ID:", b.ID)
		b.PaymentType.Refund(method)
		b.Paid = false
	} else {
		fmt.Println("The bill is not paid yet")
	}
}
