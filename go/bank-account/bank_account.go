package account

import "sync"

// Define the Account type here.
type status string

const (
	closed status = "CLOSED"
	opened status = "OPENED"
)

type Account struct {
	mu      *sync.Mutex
	balance int64
	status  status
}

func (a *Account) isClosed() bool {
	return a.status == closed
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	a := &Account{
		mu:      &sync.Mutex{},
		balance: amount,
		status:  opened,
	}
	return a
}

func (a *Account) Balance() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.isClosed() {
		return 0, false
	}
	return a.balance, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.isClosed() {
		return 0, false
	}

	if amount < 0 && a.balance < -1*amount {
		return 0, false
	}
	a.balance += amount
	return a.balance, true
}

func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.isClosed() {
		return 0, false
	}
	balance := a.balance
	a.balance = 0
	a.status = closed
	return balance, true
}
