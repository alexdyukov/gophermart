package core

import (
	"errors"
	"time"

	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

type (
	AccountWithdrawals struct {
		OperationTime time.Time
		ID            string
		OrderNumber   int
		Amount        sharedkernel.Money
	}

	Account struct {
		id              string
		user            string
		withdrawHistory []AccountWithdrawals
		balance         sharedkernel.Money
	}
)

var ErrNotEnoughFunds = errors.New("unfortunately, your account do not have enough funds")

func NewAccount(userID string) *Account {
	return &Account{
		id:              sharedkernel.NewUUID(),
		user:            userID,
		withdrawHistory: nil,
		balance:         0,
	}
}

func RestoreAccount(id, userID string, balance sharedkernel.Money) *Account {
	return &Account{
		id:              id,
		user:            userID,
		withdrawHistory: nil,
		balance:         balance,
	}
}

func (acc *Account) CurrentBalance() sharedkernel.Money {
	return acc.balance
}

func (acc *Account) WithdrawalsSum() sharedkernel.Money {
	// return cached sum or calculate on fly
	return 0
}

func (acc *Account) Add(amount sharedkernel.Money) {
	acc.balance += amount
}

// WithdrawPoints is just a representation of core model functionality behavior.
func (acc *Account) WithdrawPoints(order int, amount sharedkernel.Money, oTime time.Time) error {
	if amount > acc.balance {
		return ErrNotEnoughFunds
	}

	with := AccountWithdrawals{
		OrderNumber:   order,
		Amount:        amount,
		OperationTime: oTime,
	}

	acc.balance -= amount
	acc.withdrawHistory = append(acc.withdrawHistory, with)

	return nil
}
