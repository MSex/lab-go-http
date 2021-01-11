package stub

import (
	"io"

	"github.com/MSex/lab-go-http/app/data"
)

type Users struct {
	Data []*data.User

	OnGet        func(id data.UserId) (*data.User, error)
	OnLoadCursor func() (data.UserCursor, error)
	//TODO OnCreate
	//Todo
}

func (stub *Users) 	Create(user *data.User) (data.UserId, error) {
	return data.UserId(0), nil
}
func (stub *Users) 	ExistsLogin(login string) (bool, error) {
return false, nil
}


func (stub *Users) Read(id data.UserId) (*data.User, error) {
	if stub.OnGet != nil {
		return stub.OnGet(id)
	}

	if stub.Data == nil {
		return nil, nil
	}

	cursor, err := stub.LoadCursor()
	if err != nil {
		return nil, err
	}
	defer cursor.Close()

	for {
		user, err := cursor.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			//TODO return errors.Wrap(err, "Error getting current cursor value")
			return nil, err
		}

		if user.Id == id {
			return user, nil
		}
	}

	return nil, data.NotFoundError
}

// func (stub *Users) Create(user *User) (UserId, error)
// 	ExistsLogin(login string) (bool, error)


func (stub *Users) LoadCursor() (data.UserCursor, error) {
	if stub.OnLoadCursor != nil {
		return stub.OnLoadCursor()
	}

	if stub.Data != nil {
		return &UserCursor{Data: stub.Data}, nil
	}

	return nil, nil
}

type UserCursor struct {
	OnNext  func() (*data.User, error)
	OnClose func() error

	Data    []*data.User
	current int
}

func (stub *UserCursor) Next() (*data.User, error) {
	if stub.OnNext != nil {
		return stub.OnNext()
	}

	if stub.Data == nil {
		return nil, nil
	}
	if stub.current >= len(stub.Data) {
		return nil, io.EOF
	}

	datum := stub.Data[stub.current]
	stub.current++
	return datum, nil
}

func (stub *UserCursor) Close() error {
	if stub.OnClose != nil {
		return stub.OnClose()
	}

	return nil
}
