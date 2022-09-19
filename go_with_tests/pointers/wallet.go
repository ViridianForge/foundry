package pointers

import "fmt"

// Coin Types

type Bitcoin float64
type Ethereum float64
type Litecoin float64

// Coin Functions

func (b Bitcoin) String() string {
	return fmt.Sprintf("%.8f BTC", b)
}

func (e Ethereum) String() string {
	return fmt.Sprintf("%.8f ETH", e)
}

func (l Litecoin) String() string {
	return fmt.Sprintf("%.8f LTC", l)
}

// Wallet Type

type Wallet struct {
	btcBalance Bitcoin
	ethBalance Ethereum
	ltcBalance Litecoin
}

// Wallet Functions

func (w *Wallet) Balance() Bitcoin {
	return w.btcBalance
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.btcBalance += amount
}
