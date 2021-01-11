package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const (
	CreateTableUsers = `
		CREATE TABLE users (
    		id INTEGER PRIMARY KEY,
    		login varchar(32) NOT NULL UNIQUE,
    		name varchar(32) NOT NULL,
    		birth varchar(32) NOT NULL
		);
	`

	Insert = `
		INSERT INTO users (id, login, name, birth) VALUES (1, "renato@example.com", "Renato", "1980-11-28");
		INSERT INTO users (id, login, name, birth) VALUES (2, "guto@example.com", "Guto", "1979-12-06");
		INSERT INTO users (id, login, name, birth) VALUES (3, "mauricio@example.com", "Mauricio", "1977-05-24");
	`

)

func ProvideConn(
) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(CreateTableUsers)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(Insert)
	if err != nil {
		return nil, err
	}

	return db, nil
}


