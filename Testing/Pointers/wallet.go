package pointers

type Wallet struct {
	balance int
	name    string
}

func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount int) {
	w.balance -= amount
}

func (w Wallet) Balance() int {
	return w.balance
}
