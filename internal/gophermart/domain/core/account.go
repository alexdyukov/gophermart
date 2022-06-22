package core

import (
	"errors"
	"time"
)

type withdraw struct {
	OrderNumber int
	Amount      int
	time        int64
}

type Account struct {
	id              string
	user            string //owner
	points          int
	withdrawHistory []withdraw
}

func NewAccount(user string) Account {
	return Account{
		user: user,
	}
}

func (a *Account) CurrentPoints() int {
	return a.points
}

func (a *Account) AddPoints() { /* calculations.. */ }

func (a *Account) WithdrawPoints(order int, amount int) error {
	if amount > a.points {
		return errors.New("not enough funds")
	}
	w := withdraw{
		OrderNumber: order,
		Amount:      amount,
		time:        time.Now().Unix(),
	}
	a.points = -amount
	a.withdrawHistory = append(a.withdrawHistory, w)
	return nil
}

// ... other funcs
