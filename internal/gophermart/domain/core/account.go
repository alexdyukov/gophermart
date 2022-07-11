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
		OrderNumber   int64
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

func RestoreAccount(id, userID string, balance sharedkernel.Money, wHistory []AccountWithdrawals) *Account {
	return &Account{
		id:              id,
		user:            userID,
		withdrawHistory: wHistory,
		balance:         balance,
	}
}

func RestoreAccountWithdrawals(oTm time.Time, uid string, num int64, sum sharedkernel.Money) *AccountWithdrawals {
	return &AccountWithdrawals{
		OperationTime: oTm,
		ID:            uid,
		OrderNumber:   num,
		Amount:        sum,
	}
}

func GetSliceAccountWithdrawals(acc *Account) *[]AccountWithdrawals {
	return &acc.withdrawHistory
}

func (acc *Account) CurrentID() string {
	return acc.id
}

func (acc *Account) CurrentUserID() string {
	return acc.user
}

func (acc *Account) CurrentBalance() sharedkernel.Money {
	return acc.balance
}

func (acc *Account) WithdrawalsSum() sharedkernel.Money {
	var result sharedkernel.Money
	for _, line := range acc.withdrawHistory {
		result += line.Amount
	}
	// return cached sum or calculate on fly
	return result
}

func (acc *Account) Add(amount sharedkernel.Money) {
	acc.balance += amount
}

// WithdrawPoints is just a representation of core model functionality behavior.
func (acc *Account) WithdrawPoints(order int64, amount sharedkernel.Money) error {
	if amount > acc.balance {
		return ErrNotEnoughFunds
	}

	with := AccountWithdrawals{
		OrderNumber:   order,
		Amount:        amount,
		OperationTime: time.Now(),
	}

	acc.balance -= amount
	acc.withdrawHistory = append(acc.withdrawHistory, with)

	return nil
}
