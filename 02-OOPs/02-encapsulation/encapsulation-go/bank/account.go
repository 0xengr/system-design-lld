// package bank holds BankAccount.
// Encapsulation in Go is enforced at the package boundary:
// lowercase identifiers are invisible outside this package.

package bank

import "errors"

var (
	ErrInvalidDeposit    = errors.New("invalid deposit amount")
	ErrInvalidWithdraw   = errors.New("invalid withdraw amount")
	ErrInsufficientFunds = errors.New("insufficient balance")
)

// BankAccount keeps its balance unexported (lowercase), so code
// outside this package can't read or write it directly
type BankAccount struct {
	balance int64
}

// NewBankAccount mirrors the TS field initializer (balance = 10000).
func NewBankAccount() *BankAccount {
	return &BankAccount{
		balance: 10000,
	}
}

// defining method receivers

func (a *BankAccount) Deposit(amount int64) error {
	if amount <= 0 {
		return ErrInvalidDeposit
	}
	a.balance += amount
	return nil
}

func (a *BankAccount) Withdraw(amount int64) error {
	if amount <= 0 {
		return ErrInvalidWithdraw
	}
	if amount > a.balance {
		return ErrInsufficientFunds
	}

	a.balance -= amount
	return nil
}

// Balance is the getter. Go style drops "GET" prefix
func (a *BankAccount) Balance() int64 {
	return a.balance
}
