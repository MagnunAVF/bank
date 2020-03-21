package accounts

import (
	"errors"

	"github.com/MagnunAVF/bank/clients"
)

// SavingsAccount data
type SavingsAccount struct {
	Holder                     clients.Client
	Agency, Account, Operation int
	balance                    float64
}

// Withdraw operation
func (a *SavingsAccount) Withdraw(value float64) (float64, error) {
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
func (a *SavingsAccount) Deposit(value float64) (float64, error) {
	if value < 0 {
		return 0., errors.New("Invalid value in Deposit")
	}

	a.balance += value

	return a.balance, nil
}

// GetBalance return balance value
func (a *SavingsAccount) GetBalance() float64 {
	return a.balance
}
