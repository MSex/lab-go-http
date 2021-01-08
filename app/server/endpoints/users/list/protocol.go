package listusers

import (
	"github.com/MSex/lab-go-http/app/data"
)

type User struct {
	Id    string
	Name  string
	Login string
	Birth string
}

type ParsedRequest struct {
	Date string
}

func buildResposeUser(user *data.User) (*User, error) {
	model := &User{
		Id:    user.Id.String(),
		Name:  user.Name,
		Login: user.Login,
		Birth: user.Birth,
	}

	return model, nil
}
