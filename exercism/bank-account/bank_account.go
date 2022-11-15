package account

import (
	"math"
	"sync"
)

// Define the Account type here.
type Account struct {
	balance int64
	isOpen  bool
	mu      sync.Mutex
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	} else {
		return &Account{balance: amount, isOpen: true}
	}
}

func (a *Account) Balance() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if !a.isOpen {
		//return a.balance, false
		return 0, false
	}
	return a.balance, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	switch {
	case !a.isOpen:
		return a.balance, false
	case amount < 0 && float64(a.balance) < math.Abs(float64(amount)):
		return a.balance, false
	default:
		a.balance += amount
		return a.balance, true
	}
}

func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if !a.isOpen {
		return a.balance, false
	}
	closingBalance := a.balance
	a.balance = 0
	a.isOpen = false
	return closingBalance, true
}
