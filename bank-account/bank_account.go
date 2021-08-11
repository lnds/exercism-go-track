package account

import "sync"

type Account struct {
	open    bool
	balance int
	sync.Mutex
}

func Open(amount int) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{balance: amount, open: true}
}

func (a *Account) Deposit(amount int) (newBalance int, ok bool) {
	a.Lock()
	defer a.Unlock()
	if !a.open {
		return 0, ok
	}
	if a.balance+amount >= 0 {
		a.balance += amount
		ok = true
	}
	return a.balance, ok
}

func (a *Account) Balance() (int, bool) {
	if !a.open {
		return 0, false
	}
	return a.balance, a.open
}

func (a *Account) Close() (payout int, ok bool) {
	a.Lock()
	defer a.Unlock()
	if a.open {
		payout = a.balance
		a.open = false
		ok = true
	}
	return payout, ok
}
