package memory

import (
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
	"time"
)

type UserModel struct {
	ID       string `json:"-"`
	Login    string `json:"login"`
	Password string `json:"-"`
}

type OrderModel struct {
	UserID string              `json:"-"`
	Number string              `json:"number"`
	Status sharedkernel.Status `json:"status"`
	Sum    int                 `json:"accrual,omitempty"`
	Date   time.Time           `json:"uploaded_at"`
}

type WithdrawalModel struct {
	userID string    `json:"-"`
	Number string    `json:"number"`
	Sum    int       `json:"accrual,omitempty"`
	Date   time.Time `json:"processed_at"`
}
