package postgres

// тут создаем все таблицы, которые нам нужны
const (
	// TODO(spencer): update the CREATE DATABASE statement in the schema
	//   to pull out the database specified in the DB URL and use it instead
	//   of "photos" below.
	photosSchema = `
CREATE TABLE IF NOT EXISTS users (
 name 	  	  varchar(40)  NOT NULL,
 password	   TEXT,
PRIMARY KEY (name)
);

CREATE TABLE IF NOT EXISTS orders (
 orderNumber	varchar(15) PRIMARY KEY, 
 userName		varchar(40),
 status			varchar(10),
 accrual		numeric,
 date			timestamp,
FOREIGN KEY (userName) REFERENCES users (name) ON DELETE CASCADE
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
