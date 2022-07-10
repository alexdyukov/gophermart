package postgres

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

func (gdb *GophermartDB) SaveUserAccountTest(ctx context.Context, userId string, accrual sharedkernel.Money, withdrawal float32) error {
	fmt.Println("пытаемся сохранить баланс по аккаунту пользователя: ", userId)

	// insertSQL := `INSERT INTO user_account VALUES ($1, $2, $3, $4)
	// ON CONFLICT (uid,userId) DO UPDATE SET accrual = $3 , withdrawal = $4 ;`
	tx, err := gdb.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = gdb.saveToTableUserAccount(ctx, tx, sharedkernel.NewUUID(), userId, accrual)

	if err != nil {
		log.Printf("%v", err)
		return err
	}

	//stmt, err := tx.PrepareContext(ctx, insertSQL)
	//if err != nil {
	//	return err
	//}
	//defer stmt.Close()
	//
	//if _, err = stmt.ExecContext(ctx, sharedkernel.NewUUID(), userId, accrual, withdrawal); err != nil {
	//	return err
	//}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (gdb *GophermartDB) SaveOrderTest(ctx context.Context, userId string, numOrder int64, sum sharedkernel.Money, status sharedkernel.Status, date time.Time) error {
	exists, usrId, err := orderExists(ctx, gdb.DB, numOrder)
	if err != nil || exists {
		fmt.Println(usrId)
		return err
	}

	fmt.Println("пытаемся сохранить заказ : ", numOrder)

	trx, err := gdb.Begin()
	if err != nil {
		return err
	}

	defer trx.Rollback() // nolint:errcheck // ok

	err = gdb.saveToTableUserOrders(ctx, trx, sharedkernel.NewUUID(), userId, numOrder,
		status, sum, date)

	if err != nil {
		log.Printf("%v", err)
		return err
	}

	err = trx.Commit()

	if err != nil {
		log.Printf("%v", err)
		return err
	}

	return nil
}

func (gdb *GophermartDB) SaveWithdrawalsTest(ctx context.Context, userId string, numOrder int64, sum sharedkernel.Money, date time.Time) error {
	fmt.Println("пытаемся сохранить отоваривание по заказу : ", numOrder)

	trx, err := gdb.Begin()
	if err != nil {
		return err
	}

	defer trx.Rollback() // nolint:errcheck // ok

	err = gdb.saveToTableUserWithdrawals(ctx, trx, sharedkernel.NewUUID(), userId, numOrder,
		sum, date)
	if err != nil {
		return err
	}

	err = trx.Commit()

	if err != nil {
		log.Printf("%v", err)
		return err
	}

	return nil
}
