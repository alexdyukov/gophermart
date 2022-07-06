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
 uid TEXT NOT NULL,
 login TEXT  NOT NULL,
 passwd	   TEXT NOT NULL,
 PRIMARY KEY (uid)
);

CREATE TABLE IF NOT EXISTS orders (
 orderNumber	TEXT NOT NULL, 
 uid			TEXT,
 status			int  NOT NULL,
 accrual		numeric,
 dateAndTime	timestamp,
 PRIMARY KEY (orderNumber, uid)
);

CREATE TABLE IF NOT EXISTS withdrawals (
 orderNumber	TEXT NOT NULL,
 uid 			TEXT,
 dateAndTime	timestamp,
 amount			numeric,
 PRIMARY KEY (orderNumber, uid)
);
`
)

//FOREIGN KEY (userID) REFERENCES users (id) ON DELETE CASCADE
func InitSchema(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, photosSchema)
	return err
}

func OpenDB(dbURL string) (*sql.DB, error) {
	return sql.Open("pgx", dbURL)
}
