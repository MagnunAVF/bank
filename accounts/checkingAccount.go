package accounts

import (
	"errors"

	"github.com/MagnunAVF/bank/clients"
)

// CheckingAccount data
type CheckingAccount struct {
	Holder          clients.Client
	Agency, Account int
	Balance         float64
}

// Withdraw operation
func (a *CheckingAccount) Withdraw(value float64) (float64, error) {
	validValue := value > 0
	enoughtBalance := value <= a.Balance

	if !validValue {
		return 0., errors.New("Invalid value in Withdraw")
	}

	if !enoughtBalance {
		return 0., errors.New("Not enough funds in account")
	}

	a.Balance -= value

	return a.Balance, nil
}

// Deposit operation
func (a *CheckingAccount) Deposit(value float64) (float64, error) {
	if value < 0 {
		return 0., errors.New("Invalid value in Deposit")
	}

	a.Balance += value

	return a.Balance, nil
}
