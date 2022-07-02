package postgres

// тут создаем все таблицы, которые нам нужны
const (
	photosSchema = `
CREATE TABLE IF NOT EXISTS users (
 id TEXT,
 varchar(40)  NOT NULL,
 password	   TEXT,
 PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS orders (
 orderNumber	varchar(15) PRIMARY KEY, 
 userID		varchar(40),
 status			int,
 accrual		numeric,
 dateAndTime			timestamp,
FOREIGN KEY (userID) REFERENCES users (name) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS withdrawals (
 orderNumber	varchar(15) PRIMARY KEY,
 userName 		varchar(40),
 date			timestamp,
 amount			numeric,
FOREIGN KEY (userName) REFERENCES users (name) ON DELETE CASCADE
);
`
)
