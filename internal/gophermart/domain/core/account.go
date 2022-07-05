package core

import (
	"errors"
	"time"
)

// Делала BeOl

type withdraw struct {
	OrderNumber string
	Amount      float32
	Time        time.Time
}

type Account struct {
	Id              string
	User            string //owner
	Points          float32
	WithdrawHistory []withdraw
}

func NewAccount(user string) Account {
	return Account{
		User: user,
	}
}

func (a *Account) CurrentPoints() float32 {
	return a.Points
}

func (a *Account) AddPoints() { /* calculations.. */ }

func (a *Account) WithdrawPoints(order string, amount float32) error {
	if amount > a.Points {
		return errors.New("not enough funds")
	}
	w := withdraw{
		OrderNumber: order,
		Amount:      amount,
		Time:        time.Now(),
	}
	a.Points -= amount //  тут меняла потому что так кажется логичным. Но вообще не обязательно, надо разбираться
	a.WithdrawHistory = append(a.WithdrawHistory, w)
	return nil
}

// ... other funcs
