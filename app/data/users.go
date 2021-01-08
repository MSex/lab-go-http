package data

import (
	"fmt"
	"strconv"
)

type UserId int32

type User struct {
	Id   UserId
	Name string
	Login string
	Birth string 
}

type UserCursor interface {
	Next() (*User, error)
	Close() error
}

type Users interface {
	Get(id UserId) (*User, error)
	LoadCursor() (UserCursor, error)
	Create(user *User) (UserId, error)
	ExistsLogin(login string) (bool, error)

}

func UserIdFromString(id string) (UserId, error) {
	if id == "" {
		return 0, nil
	}

	int, err := strconv.Atoi(id)
	return UserId(int), err
}

func (id *UserId) String() string {
	if id == nil {
		return ""
	}

	return fmt.Sprintf("%d", *id)
}

func (id *UserId) Int32() int32 {
	if id == nil {
		return 0
	}

	return int32(*id)
}
