package sqlite

import (
	"database/sql"
)

const (
	CreateTableUsers = `
		CREATE TABLE users (
    		id varchar(3) NOT NULL PRIMARY KEY,
    		login varchar(32) NOT NULL UNIQUE,
    		name varchar(32) NOT NULL UNIQUE,
    		birth varchar(32) NOT NULL UNIQUE
		);
	`

	Insert = `
		INSERT INTO users (login, name, birth) VALUES ("renato@example.com", "Renato", "1980-11-28");
		INSERT INTO users (login, name, birth) VALUES ("guto@example.com", "Guto", "1979-12-06");
		INSERT INTO users (login, name, birth) VALUES ("mauricio@example.com", "Mauricio", "1977-05-24");
	`

)

func Open(
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


