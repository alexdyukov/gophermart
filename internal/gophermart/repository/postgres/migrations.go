package postgres

import (
	"context"
	"database/sql"
	_ "github.com/jackc/pgx/stdlib"
)

// тут создаем все таблицы, которые нам нужны
const (
	photosSchema = `
CREATE TABLE IF NOT EXISTS users (
 id varchar(40),
 uName varchar(40)  NOT NULL,
 uPassword	   TEXT,
 PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS orders (
 orderNumber	varchar(15) PRIMARY KEY, 
 userID		varchar(40),
 status			int,
 accrual		numeric,
 dateAndTime	timestamp,
FOREIGN KEY (userID) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS withdrawals (
 orderNumber	varchar(15) PRIMARY KEY,
 userID 		varchar(40),
 dateAndTime	timestamp,
 amount			numeric,
FOREIGN KEY (userID) REFERENCES users (id) ON DELETE CASCADE
);
`
)

func InitSchema(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, photosSchema)
	return err
}

func OpenDB(dbURL string) (*sql.DB, error) {
	return sql.Open("pgx", dbURL)
}
