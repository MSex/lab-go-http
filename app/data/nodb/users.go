package nodb

import (
	"github.com/MSex/lab-go-http/app/data"
	"github.com/MSex/lab-go-http/app/data/stub"
)

func ProvideUsers() data.Users {
	return &stub.Users{
		Data: []*data.User{
			{Id: 1, Name: "Guto", Login: "guto@example.com", Birth: "1978-12-06"},
			{Id: 2, Name: "Renato", Login: "renato@example.com", Birth: "1979-11-28"},
			{Id: 3, Name: "MSex", Login: "mauricio@example.com", Birth: "1977-05-24"},
		},
	}
}
