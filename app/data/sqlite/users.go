package sqlite

import (
	"database/sql"
	"io"

	"github.com/MSex/lab-go-http/app/data"
	"github.com/pkg/errors"
)

type Users struct {
	db *sql.DB
}
 
func ProvideUsers(db *sql.DB) data.Users {
	return &Users{db: db}
}

func (users *Users) Create(user *data.User) (data.UserId, error) {
	queryString := `
		INSERT INTO users
		(login, name, birth)
		VALUES (?,?,?);
	`

	params := []interface{}{
		user.Login,
		user.Name,
		user.Birth,
	}


	res,  err := users.db.Exec(queryString, params...)
	if err != nil {
		return 0, errors.Wrap(err, "Error inserting into database")
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, errors.Wrap(err, "Error getting last insert id")
	}

	return data.UserId(id), nil
}

func (users *Users) ExistsLogin(login string) (bool, error) {
	queryString := ` 
	SELECT EXISTS (
		SELECT * FROM users 
		WHERE login = ?
	)`

	var exists bool

	err := users.db.QueryRow(queryString, login).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (users *Users) Read(id data.UserId) (*data.User, error) {
	queryString := `
	SELECT 
		id, 
		login, 
		name, 
		birth
	FROM users
	WHERE id = ?
	`

	rows, err := users.db.Query(queryString, id)
	if err != nil {
		return nil, errors.Wrap(err, "Error executing query")
	}
	defer rows.Close()

	if rows.Next() {
		return buildUser(rows)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "Failed to iterate over query results")
	}
	return nil, data.NotFoundError
}

func (users *Users) LoadCursor() (data.UserCursor, error) {
	queryString := `
	SELECT 
		id, 
		login, 
		name, 
		birth 
	FROM users
	`

	rows, err := users.db.Query(queryString)
	if err != nil {
		return nil, errors.Wrap(err, "Error executing query")
	}
	return &userCursor{rows: rows}, nil
}


func (users *Users) Update(userId data.UserId, user *data.User) error {
	queryString := `
	UPDATE users 
	SET 
		name = ?,
		birth = ?
	WHERE id = ?
	`

	params := []interface{}{
		user.Name,
		user.Birth,
		userId,
	}

	_, err := users.db.Exec(queryString, params...)
	if err != nil {
		return errors.Wrap(err, "Error updating database")
	}

	return nil
}

func (users *Users) Delete(userId data.UserId) error {
	queryString := `
		DELETE FROM users
		WHERE id = ?
	`

	_, err := users.db.Exec(queryString, userId)
	if err != nil {
		return errors.Wrap(err, "Error deleting user from database")
	}

	return nil
}


type userCursor struct {
	rows *sql.Rows
}

func (cursor *userCursor) Next() (*data.User, error) {
	if cursor.rows.Next() {
		return buildUser(cursor.rows)
	}
	if err := cursor.rows.Err(); err != nil {
		return nil, errors.Wrap(err, "Failed to iterate over query results")
	}
	return nil, io.EOF
}

func buildUser(rows *sql.Rows) (*data.User, error) {
	var (
		id               int32
		login         string
		name        string
		birth       string
	)

	err := rows.Scan(
		&id,
		&login,
		&name,
		&birth,
	)
	if err != nil {
		return nil, err
	}
	user := &data.User{
		Id:         data.UserId(id),
		Login:  login,
		Name: name,
		Birth: birth,
	}

	return user, nil
}

func (cursor *userCursor) Close() error {
	return cursor.rows.Close()
}
