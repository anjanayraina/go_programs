package pointers

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Testing Deposit", func(t *testing.T) {
		wallet := Wallet{balance: 0, name: "Anjanay Raina"}
		wallet.Deposit(100)
		if wallet.Balance() != 100 {
			t.Errorf("The balance is wrong ")
		}

		wallet.Withdraw(90)
		if wallet.Balance() != 10 {
			t.Errorf("The balance is wrong ")
		}
	})
}
