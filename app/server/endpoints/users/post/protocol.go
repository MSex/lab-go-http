package postusers

import (
	"github.com/MSex/lab-go-http/app/data"
)

type User struct {
	Name  string
	Login string
	Birth string
}

type RequestBody struct {
	User User
}

type ParsedRequest struct {
	User User
}

type ResponseBody struct {
	Id string
}


func buildUser(user *User) (*data.User, error) {
	model := &data.User{
		Name:  user.Name,
		Login: user.Login,
		Birth: user.Birth,
	}

	return model, nil
}
