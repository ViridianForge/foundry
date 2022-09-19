package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	wallet := Wallet{}

	wallet.Deposit(10.0)

	got := wallet.Balance()
	expected := Bitcoin(10.0)

	if got != expected {
		t.Errorf("got %s, expected %s", got, expected)
	}
}
