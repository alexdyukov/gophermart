package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

func (gdb *GophermartDB) SaveUserAccountTest(ctx context.Context, userId string, accrual float32, withdrawal float32) error {

	insertSQL := `INSERT INTO user_account VALUES ($1, $2, $3, $4)
	ON CONFLICT (uid,userId) DO UPDATE SET accrual = $3 , withdrawal = $4 ;`
	tx, err := gdb.Begin()

	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, insertSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.ExecContext(ctx, sharedkernel.NewUUID(), userId, accrual, withdrawal); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil

}

func (gdb *GophermartDB) SaveOrderTest(ctx context.Context, userId string, numOrder int, sum float32, status sharedkernel.Status, date time.Time) error {
	exists, err := orderExists(ctx, gdb.DB, numOrder)
	if err != nil || exists {
		return err
	}

	fmt.Println("пытаемся сохранить заказ : ", numOrder)
	const insertSQL = `
	INSERT INTO user_orders VALUES ($1, $2, $3, $4, $5, $6);
	`

	_, err = gdb.ExecContext(ctx, insertSQL, sharedkernel.NewUUID(), numOrder, userId, status, sum, date)

	return err
}

func (gdb *GophermartDB) SaveWithdrawalsTest(ctx context.Context, userId string, numOrder int, sum float32, date time.Time) error {
	exists, err := withdrawalExists(ctx, gdb.DB, numOrder)
	if err != nil || exists {
		return err
	}

	fmt.Println("пытаемся сохранить отоваривание по заказу : ", numOrder)
	const insertSQL = `
	INSERT INTO user_withdrawals VALUES ($1, $2, $3, $4, $5);
	`

	_, err = gdb.ExecContext(ctx, insertSQL, sharedkernel.NewUUID(), numOrder, userId, sum, date)

	return err
}

// userExists looks up a user by ID.
func orderExists(ctx context.Context, db *sql.DB, orderNumber int) (bool, error) {
	var id int
	const selectSQL = `
SELECT orderNumber FROM user_orders WHERE orderNumber = $1;
`
	err := db.QueryRowContext(ctx, selectSQL, orderNumber).Scan(&id)
	switch err {
	case sql.ErrNoRows:
		return false, nil
	case nil:
		return true, nil
	default:
		return false, err
	}
}

// userExists looks up a user by ID.
func withdrawalExists(ctx context.Context, db *sql.DB, orderNumber int) (bool, error) {
	var id int
	const selectSQL = `
SELECT orderNumber FROM user_withdrawals WHERE orderNumber = $1;
`
	err := db.QueryRowContext(ctx, selectSQL, orderNumber).Scan(&id)
	switch err {
	case sql.ErrNoRows:
		return false, nil
	case nil:
		return true, nil
	default:
		return false, err
	}
}
