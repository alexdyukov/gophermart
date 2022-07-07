package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
	"time"
)

type GophermartDB struct {
	*sql.DB
}

func NewGophermartDB(conn *sql.DB) (*GophermartDB, error) {
	dataBase := GophermartDB{
		conn,
	}

	err := dataBase.createTablesIfNotExist()
	if err != nil {
		return nil, err
	}

	return &dataBase, nil
}

func (gdb *GophermartDB) GetOrdersByUser(ctx context.Context, uid string) ([]core.UserOrderNumber, error) {
	result := make([]core.UserOrderNumber, 0)

	query := `
	SELECT
	uid,
	orderNumber,
	status,
	accrual,
	dateAndTime
	FROM user_orders
	WHERE userID = $1
	`
	rows, err := gdb.QueryContext(ctx, query, uid)
	// only one cuddle assignment allowed before if statement for linter
	if err != nil {
		return result, err
	}
	defer rows.Close()

	ord := core.UserOrderNumber{}

	for rows.Next() {
		err = rows.Scan(&ord.Id, &ord.Number, &ord.Status, &ord.Accrual, &ord.DateAndTime)
		if err != nil {
			return nil, err
		}

		result = append(result, ord)
	}

	return result, nil
}

func (gdb *GophermartDB) GetAccountByID(ctx context.Context, id string) (core.Account, error) {
	// work with db
	return core.Account{}, nil
}

func (gdb *GophermartDB) GetBalance(_ context.Context, user string) (float32, float32, error) {

	// work with db
	return 0, 0, nil
}

func (gdb *GophermartDB) SaveOrderNumber(core.UserOrderNumber) error {

	// work with db
	return nil
}

func (gdb *GophermartDB) SaveAccount(ctx context.Context, acc core.Account) error {

	//exists, err := userExists(ctx,p.db , userID)
	//if err != nil || exists {
	//	return err
	//}
	//	const insertSQL = `
	//INSERT INTO users VALUES ($1, 0, 0, $2, $3);
	//`
	//	const minNameLen = 1
	//	const maxNameLen = 30
	//	const minAddrLen = 20
	//	const maxAddrLen = 100
	//	_, err = p.db.ExecContext(ctx, insertSQL, acc.User, name, addr)
	//	return err
	return nil
}

func (gdb *GophermartDB) SaveOrderTest(ctx context.Context, userId string, numOrder int, sum float32) error {

	exists, err := orderExists(ctx, gdb.DB, numOrder)
	if err != nil || exists {
		return err
	}

	fmt.Println("пытаемся сохранить заказ : ", numOrder)
	const insertSQL = `
	INSERT INTO user_orders VALUES ($1, $2, $3, $4, $5, $6);
	`

	_, err = gdb.ExecContext(ctx, insertSQL, sharedkernel.NewUUID(), numOrder, userId, sharedkernel.NEW, sum, time.Now())

	return err

}

func (gdb *GophermartDB) SaveUser(ctx context.Context, name string, passw string, userID string) error {

	//

	exists, err := userExists(ctx, gdb.DB, userID)
	if err != nil || exists {
		return err
	}
	fmt.Println("пытаемся сохранить пользователя: ", name, userID)

	const insertSQL = `
	INSERT INTO users VALUES ($1, $2, $3);
	`

	_, err = gdb.ExecContext(ctx, insertSQL, userID, name, passw)
	if err != nil {
		fmt.Println("какая-то ошибка при сохранении юзера в бд")
		return err
	}

	return nil
}

// userExists looks up a user by ID.
func userExists(ctx context.Context, db *sql.DB, userID string) (bool, error) {
	var id string
	const selectSQL = `
SELECT id FROM users WHERE id = $1;
`
	row := db.QueryRowContext(ctx, selectSQL, userID)

	err := row.Scan(&id)
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

func (gdb *GophermartDB) createTablesIfNotExist() error {
	_, err := gdb.Exec(`CREATE TABLE IF NOT EXISTS public.user_orders (
    											uid TEXT NOT NULL,
     											orderNumber	bigint NOT NULL, 
												userID TEXT,
												status INT  NOT NULL,
												accrual		numeric,
												dateAndTime	timestamp,
												PRIMARY KEY (uid)
												);

								CREATE TABLE IF NOT EXISTS public.user_withdrawals (
								    			uid TEXT NOT NULL,
     											orderNumber	bigint NOT NULL, 
												userID TEXT,
												amount		numeric,
												dateAndTime	timestamp,
												PRIMARY KEY (uid)
												);
												`)
	if err != nil {
		return err // nolint:wrapcheck // ok
	}

	return nil
}
