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
 id TEXT,
 uName TEXT  NOT NULL,
 uPassword	   TEXT,
 PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS orders (
 orderNumber	TEXT PRIMARY KEY, 
 userID			TEXT PRIMARY KEY,
 status			int,
 accrual		numeric,
 dateAndTime	timestamp,
FOREIGN KEY (userID) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS withdrawals (
 orderNumber	TEXT PRIMARY KEY,
 userID 		TEXT,
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
