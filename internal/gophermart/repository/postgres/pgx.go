package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
	"time"
)

type GophermartStore struct {
	db *sql.DB
}

func NewGophermartStore(datb *sql.DB) *GophermartStore {

	return &GophermartStore{db: datb}
}

func (p *GophermartStore) GetOrdersByUser(ctx context.Context, user string) ([]core.OrderNumber, error) {

	selectSQL := `
	SELECT 
	 *
	FROM orders;
	`

	//selectSQL := `
	//SELECT
	// orderNumber,
	// status,
	// accrual,
	// dateAndTime
	//FROM orders
	//WHERE userID=$1 ;
	//`
	rez := make([]core.OrderNumber, 0)

	rows, err := p.db.QueryContext(ctx, selectSQL)
	if err != nil {
		fmt.Println("возникла ошибка в процессе получения списка заказов ", err)
		return rez, err
	}

	defer rows.Close()

	for rows.Next() {
		var ord core.OrderNumber
		err = rows.Scan(&ord.Number, ord.Status, &ord.Accrual, &ord.Data)
		if err != nil {
			return nil, err
		}
		rez = append(rez, ord)
	}
	return rez, nil
}

func (p *GophermartStore) GetAccountByID(ctx context.Context, id string) (core.Account, error) {
	// work with db
	return core.Account{}, nil
}

func (p *GophermartStore) GetBalance(_ context.Context, user string) (float32, float32, error) {

	// work with db
	return 0, 0, nil
}

func (p *GophermartStore) SaveOrderNumber(core.OrderNumber) error {

	// work with db
	return nil
}

func (p *GophermartStore) SaveAccount(ctx context.Context, acc core.Account) error {

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

func (p *GophermartStore) SaveOrderTest(ctx context.Context, userId string, numOrder string, sum float32) error {
	fmt.Println("пытаемся сохранить заказ : ", numOrder)
	exists, err := orderExists(ctx, p.db, numOrder)
	if err != nil || exists {
		return err
	}
	const insertSQL = `
	INSERT INTO orders VALUES ($1, $2, $3, $4);
	`
	_, err = p.db.ExecContext(ctx, insertSQL, userId, sharedkernel.NEW, sum, time.Now())
	return err
	return nil
}

func (p *GophermartStore) SaveUser(ctx context.Context, name string, passw string, userID string) error {

	//
	fmt.Println("пытаемся сохранить пользователя: ", name, userID)
	exists, err := userExists(ctx, p.db, userID)
	if err != nil || exists {
		return err
	}
	const insertSQL = `
	INSERT INTO users VALUES ($1, $2, $3);
	`

	_, err = p.db.ExecContext(ctx, insertSQL, userID, name, passw)
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
func orderExists(ctx context.Context, db *sql.DB, userID string) (bool, error) {
	var id int
	const selectSQL = `
SELECT id FROM users WHERE id = $1;
`
	err := db.QueryRowContext(ctx, selectSQL, userID).Scan(&id)
	switch err {
	case sql.ErrNoRows:
		return false, nil
	case nil:
		return true, nil
	default:
		return false, err
	}
}
