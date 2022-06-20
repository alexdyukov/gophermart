package storage

import "time"

const (
	StNew        = "NEW"
	StProcessing = "PROCESSING"
	StInvalid    = "INVALID"
	StProcessed  = "PROCESSED"
)

type UsersModel struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password,omitempty"`
}

type OrdersModel struct {
	UsersModel `json:"-"`
	Number     int       `json:"number"`
	Status     string    `json:"status"`
	Sum        int       `json:"accrual,omitempty"`
	Date       time.Time `json:"uploaded_at"`
}

type WithdrawalsModel struct {
	UsersModel `json:"-"`
	Number     int       `json:"number"`
	Sum        int       `json:"accrual,omitempty"`
	Date       time.Time `json:"processed_at"`
}
