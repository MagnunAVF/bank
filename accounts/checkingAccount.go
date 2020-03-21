package accounts

import (
	"errors"

	"github.com/MagnunAVF/bank/clients"
)

// CheckingAccount data
type CheckingAccount struct {
	Holder          clients.Client
	Agency, Account int
	balance         float64
}

// Withdraw operation
func (a *CheckingAccount) Withdraw(value float64) (float64, error) {
	validValue := value > 0
	enoughBalance := value <= a.balance

	if !validValue {
		return 0., errors.New("Invalid value in Withdraw")
	}

	if !enoughBalance {
		return 0., errors.New("Not enough funds in account")
	}

	a.balance -= value

	return a.balance, nil
}

// Deposit operation
func (a *CheckingAccount) Deposit(value float64) (float64, error) {
	if value < 0 {
		return 0., errors.New("Invalid value in Deposit")
	}

	a.balance += value

	return a.balance, nil
}

// Transfer a value between accounts
func (a *CheckingAccount) Transfer(value float64, destinyAccount *CheckingAccount) error {
	validValue := value > 0
	enoughBalance := value <= a.balance

	if !validValue {
		return errors.New("Invalid value in Withdraw")
	}

	if !enoughBalance {
		return errors.New("Not enough funds in account")
	}

	a.balance -= value
	destinyAccount.Deposit(value)

	return nil
}

// GetBalance return balance value
func (a *CheckingAccount) GetBalance() float64 {
	return a.balance
}
