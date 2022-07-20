package postgres

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strconv"

	"github.com/alexdyukov/gophermart/internal/accrual/domain/core"
	"github.com/alexdyukov/gophermart/internal/accrual/domain/usecase"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
	"github.com/jackc/pgconn"
)

type AccrualDB struct {
	*sql.DB
}

const (
	pgxErrorRecordDuplicate = "23505"
)

func NewAccrualDB(conn *sql.DB) (*AccrualDB, error) {
	accrualDB := AccrualDB{
		conn,
	}

	err := accrualDB.createUserTableIfNotExist()
	if err != nil {
		return nil, err
	}

	return &accrualDB, nil
}

// nolint:funlen // ok
func (accdb *AccrualDB) SaveOrderReceipt( // nolint:gocognit,cyclop // ok
	ctx context.Context,
	orderReceipt *core.OrderReceipt,
) error { // nolint:whitespace // ok
	transaction, err := accdb.Begin()
	if err != nil {
		return err
	}

	defer func() {
		err = transaction.Rollback()
		if err == nil {
			log.Println(err)
		}
	}()

	stmt, err := transaction.Prepare(`INSERT INTO public.accrual_orders (order_number_id, accrual, status)
									  VALUES ($1,$2,$3)`)
	if err != nil {
		return err
	}

	defer func() {
		err = stmt.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	stmtProducts, err := transaction.Prepare(`INSERT INTO public.accrual_products (order_number, description, price)
VALUES ($1,$2,$3)`)
	if err != nil {
		return err
	}
	defer stmtProducts.Close()

	_, err = stmt.ExecContext(ctx, orderReceipt.OrderNumber, orderReceipt.Accrual, orderReceipt.Status.String())
	if err != nil {
		var pgError *pgconn.PgError
		if ok := errors.As(err, &pgError); ok {
			if pgError.Code == pgxErrorRecordDuplicate {
				return usecase.ErrOrderAlreadyExist
			}

			return err
		}

		return err
	}

	for _, product := range orderReceipt.Goods {
		if _, err = stmtProducts.Exec(orderReceipt.OrderNumber, product.Description, product.Price); err != nil {
			return err
		}
	}

	err = transaction.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (accdb *AccrualDB) SaveRewardMechanic(ctx context.Context, reward *core.Reward) error {
	rewardStmt, err := accdb.Prepare(`INSERT INTO public.accrual_rewards (match, reward_amount, reward_type)
 									VALUES ($1,$2,$3)`)
	if err != nil {
		return err
	}

	defer func() {
		err = rewardStmt.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	_, err = rewardStmt.ExecContext(ctx, reward.Match(), reward.RewardPoints(), reward.RewardType())
	if err != nil {
		var pgError *pgconn.PgError
		if ok := errors.As(err, &pgError); ok {
			if pgError.Code == pgxErrorRecordDuplicate {
				return usecase.ErrRewardAlreadyExists
			}

			return err
		}

		return err
	}

	return nil
}

func (accdb *AccrualDB) GetOrderByNumber(ctx context.Context, number int64) (*core.OrderReceipt, error) {
	orderStmt, err := accdb.PrepareContext(ctx, `SELECT
											   order_number_id,
											   accrual,
											   status
											  FROM public.accrual_orders
											  WHERE order_number_id = $1
											  LIMIT 1`)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = orderStmt.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	orderReceipt := core.OrderReceipt{}

	var statusStr string

	err = orderStmt.QueryRowContext(ctx, number).Scan(&orderReceipt.OrderNumber, &orderReceipt.Accrual, &statusStr)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, usecase.ErrOrderReceiptNotExist
		}

		return nil, err
	}

	statusStrQ := strconv.Quote(statusStr)

	err = orderReceipt.Status.UnmarshalJSON([]byte(statusStrQ))
	if err != nil {
		return nil, err
	}

	return &orderReceipt, nil
}

// nolint:gocognit // ok
func (accdb *AccrualDB) GetOrderByNumberWithGoods( // nolint:funlen,cyclop // ok
	ctx context.Context, number int64,
) (*core.OrderReceipt, error) { // nolint:whitespace // ok
	transaction, err := accdb.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		err = transaction.Rollback()
		if err != nil {
			log.Println(err)
		}
	}()

	orderStmt, err := transaction.PrepareContext(ctx, `SELECT
											   order_number_id,
											   accrual,
											   status
											  FROM public.accrual_orders
											  WHERE order_number_id = $1
											  LIMIT 1`)

	defer func() {
		err = orderStmt.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	orderReceipt := core.OrderReceipt{}

	var statusStr string

	err = orderStmt.QueryRowContext(ctx, number).Scan(&orderReceipt.OrderNumber, &orderReceipt.Accrual, &statusStr)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, usecase.ErrOrderReceiptNotExist
		}

		return nil, err
	}

	statusStrQ := strconv.Quote(statusStr)

	err = orderReceipt.Status.UnmarshalJSON([]byte(statusStrQ))
	if err != nil {
		return nil, err
	}

	productsStmt, err := transaction.PrepareContext(ctx, `SELECT description, price
															FROM accrual_products WHERE order_number = 1$ `)
	defer func() {
		err = orderStmt.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	productRows, err := productsStmt.QueryContext(ctx, orderReceipt.OrderNumber)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

		return nil, err
	}
	defer productRows.Close()

	var products []core.Product

	for productRows.Next() {
		product := core.Product{}

		err = productRows.Scan(product.Description, product.Price)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	err = productRows.Err()
	if err != nil {
		return nil, err
	}

	orderReceipt.Goods = products

	err = transaction.Commit()
	if err != nil {
		return nil, err
	}

	return &orderReceipt, nil
}

func (accdb *AccrualDB) FindAllRewardMechanicsByTokens( // nolint:funlen,cyclop,gocognit // ok
	ctx context.Context, tokens ...string,
) (map[string]core.Reward, error) { // nolint:whitespace // ok
	query := `SELECT match, reward_amount, reward_type FROM public.accrual_rewards WHERE `

	for key := range tokens {
		if key == 0 {
			query = query + "match = $" + strconv.Itoa(key+1)

			continue
		}

		query = query + " OR match  = $" + strconv.Itoa(key+1)
	}

	rewardsStmt, err := accdb.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = rewardsStmt.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	values := make([]interface{}, 0)
	for _, v := range tokens {
		values = append(values, v)
	}

	rows, err := rewardsStmt.QueryContext(ctx, values...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, usecase.ErrNoRewards
		}

		return nil, err
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	rewards := make(map[string]core.Reward, len(tokens))

	for rows.Next() {
		rewardAmount := sharedkernel.Money(0)
		rewardType := ""
		match := ""

		err = rows.Scan(&match, &rewardAmount, &rewardType)
		if err != nil {
			return nil, err
		}

		reward := core.RestoreReward(match, rewardAmount, rewardType)
		rewards[reward.Match()] = *reward
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return rewards, nil
}

func (accdb *AccrualDB) UpdateReceiptOrderState(ctx context.Context, orderReceipt *core.OrderReceipt) error {
	query := `UPDATE accrual_orders SET accrual = $1, status = $2 WHERE order_number_id = $3`

	stmt, err := accdb.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	defer func() {
		err = stmt.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	_, err = stmt.ExecContext(ctx, orderReceipt.Accrual, orderReceipt.Status, orderReceipt.OrderNumber)
	if err != nil {
		return err
	}

	return nil
}

func (accdb *AccrualDB) createUserTableIfNotExist() error {
	_, err := accdb.Exec(`CREATE TABLE IF NOT EXISTS public.accrual_orders (
	                       order_number_id BIGINT NOT NULL,
						   accrual NUMERIC(12,2) NOT NULL,
						   status TEXT NOT NULL,
						   CONSTRAINT order_number_pk_constraint PRIMARY KEY (order_number_id));

                       CREATE TABLE IF NOT EXISTS public.accrual_products (
                           id SERIAL PRIMARY KEY,
                           order_number BIGINT NOT NULL,
						   description TEXT NOT NULL,
						   price NUMERIC(12,2) NOT NULL,
                           FOREIGN KEY (order_number) REFERENCES public.accrual_orders (order_number_id));
                           
                       CREATE TABLE IF NOT EXISTS public.accrual_rewards (
	                       match TEXT NOT NULL,
	                       reward_amount NUMERIC(12,2) NOT NULL,
						   reward_type TEXT NOT NULL,
						   CONSTRAINT match_constraint PRIMARY KEY (match));
					   `)
	if err != nil {
		return err
	}

	return nil
}
