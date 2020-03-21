package accounts

import "github.com/MagnunAVF/bank/clients"

type CheckingAccount struct {
	Holder          clients.Client
	Agency, Account int
	balance         float64
}
