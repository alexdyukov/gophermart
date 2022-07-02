package memory

import (
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
	"time"
)

type UserModel struct {
	ID       string
	Login    string
	Password string
}

type OrderModel struct {
	UserID string
	Number string
	Status sharedkernel.Status
	Sum    float32
	Date   time.Time
}

type WithdrawalModel struct {
	UserID string
	Number string
	Sum    float32
	Date   time.Time
}
