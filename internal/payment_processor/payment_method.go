package paymentprocessor

type PaymentMethod interface {
	getBalance() float64
	deposit(amount float64) bool
	withdraw(amount float64) bool
	isCard() bool
	getCardNumber() string
}

type Card struct {
	CardNumber string
	Balance    float64
}

func (c *Card) getBalance() float64 {
	return c.Balance
}

func (c *Card) deposit(amount float64) bool {
	if amount > 0 {
		c.Balance += amount
		return true
	}
	return false
}

func (c *Card) withdraw(amount float64) bool {
	if amount > 0 && amount <= c.Balance {
		c.Balance -= amount
		return true
	}
	return false
}

func (c *Card) isCard() bool {
	return true
}

func (c *Card) getCardNumber() string {
	return c.CardNumber
}

type BankAccount struct {
	accountNumber string
	Balance       float64
}

func (ba *BankAccount) getBalance() float64 {
	return ba.Balance
}

func (ba *BankAccount) deposit(amount float64) bool {
	if amount > 0 {
		ba.Balance += amount
		return true
	}
	return false
}

func (ba *BankAccount) withdraw(amount float64) bool {
	if amount > 0 && amount <= ba.Balance {
		ba.Balance -= amount
		return true
	}
	return false
}

func (ba *BankAccount) isCard() bool {
	return false
}

func (ba *BankAccount) getCardNumber() string {
	return ""
}

type Cash struct {
	Balance float64
}

func (c *Cash) getBalance() float64 {
	return c.Balance
}

func (c *Cash) deposit(amount float64) bool {
	if amount > 0 {
		c.Balance += amount
		return true
	}
	return false
}

func (c *Cash) withdraw(amount float64) bool {
	if amount > 0 && amount <= c.Balance {
		c.Balance -= amount
		return true
	}
	return false
}

func (c *Cash) isCard() bool {
	return false
}

func (c *Cash) getCardNumber() string {
	return ""
}
