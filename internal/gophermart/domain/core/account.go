package core

import (
	"errors"
	"time"

	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

// Pay attention this is the first iteration (a sort of Draft)
// of the core structure, so no warranty it is correct or will not change.
type (
	withdraw struct {
		OrderNumber int
		Amount      int
		time        int64
	}

	Account struct {
		id              string
		user            string
		withdrawHistory []withdraw
		balance         int
	}
)

var ErrNotEnoughFunds = errors.New("account do not have enough funds")

func NewAccount(userID string) *Account {
	return &Account{ // nolint:exhaustivestruct // ok.
		id:   sharedkernel.NewUUID(),
		user: userID,
	}
}

func RestoreAccount(id, userID string, balance int) *Account {
	return &Account{
		id:              id,
		user:            userID,
		withdrawHistory: nil,
		balance:         balance,
	}
}

func (acc *Account) ShowBalance() int {
	return acc.balance
}

func (acc *Account) Add(amount int) {
	acc.balance += amount
}

// WithdrawPoints is just a representation of core model functionality (an example of core model behavior).
func (acc *Account) WithdrawPoints(order, amount int) error {
	if amount > acc.balance {
		return ErrNotEnoughFunds
	}

	with := withdraw{
		OrderNumber: order,
		Amount:      amount,
		time:        time.Now().Unix(),
	}

	acc.balance = -amount
	acc.withdrawHistory = append(acc.withdrawHistory, with)

	return nil
}
