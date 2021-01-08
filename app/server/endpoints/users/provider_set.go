package users

import (
	getusers "github.com/MSex/lab-go-http/app/server/endpoints/users/get"
	listusers "github.com/MSex/lab-go-http/app/server/endpoints/users/list"
	postusers "github.com/MSex/lab-go-http/app/server/endpoints/users/post"

	"github.com/google/wire"
)

// ProviderSet represents the injection module
var ProviderSet = wire.NewSet(
	listusers.ProvideHandler,
	postusers.ProvideHandler,
	getusers.ProvideHandler,
)
