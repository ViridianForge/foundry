package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit BTC", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10.0)
		assertBtcBalance(t, wallet, Bitcoin(10.0))
	})

	t.Run("Withdraw BTC", func(t *testing.T) {
		wallet := Wallet{btcBalance: 10.0}
		err := wallet.Withdraw(Bitcoin(7.5))
		assertNoError(t, err)
		assertBtcBalance(t, wallet, Bitcoin(2.5))
	})

	t.Run("Withdraw insufficient BTC", func(t *testing.T) {
		wallet := Wallet{btcBalance: 15.0}
		err := wallet.Withdraw(Bitcoin(20.0))
		assertError(t, err, ErrInsufficientFunds)
		assertBtcBalance(t, wallet, Bitcoin(15.0))
	})
}

func assertBtcBalance(t testing.TB, wallet Wallet, expected Bitcoin) {
	t.Helper()
	got := wallet.btcBalance

	if got != expected {
		t.Errorf("got %s, expected %s", got, expected)
	}
}

func assertError(t testing.TB, got error, expected error) {
	t.Helper()
	if got == nil {
		t.Error("an error was expected, but not observed")
	}
	if got != expected {
		t.Errorf("got %q, expected %q", got, expected)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Error("no error was expected, but one was observed")
	}
}
