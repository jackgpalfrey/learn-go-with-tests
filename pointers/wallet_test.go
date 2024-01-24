package pointers

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	t.Run("deposit", func(t *testing.T){
		wallet.Deposit(10)

		got := wallet.Balance()
		want := Pounds(10)

		if got != want {
			t.Errorf("expected %s, but got %s", want, got)
		}
	})

	t.Run("withdraw", func(t *testing.T){
		wallet.Withdraw(6)

		got := wallet.Balance()
		want := Pounds(4)

		if got != want {
			t.Errorf("expected %s, but got %s", want, got)
		}
	})

	t.Run("withdraw insufficent funds", func(t *testing.T){
		err := wallet.Withdraw(6)

		if err == nil {
			t.Errorf("wanted error but didn't get one")
		}
	})
}
