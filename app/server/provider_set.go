package server

import (
	"github.com/MSex/lab-go-http/app/server/endpoints/users"
	"github.com/google/wire"
)

// ProviderSet represents the injection module
var ProviderSet = wire.NewSet(
	users.ProviderSet,
	

	ProvideRouter, 
)
