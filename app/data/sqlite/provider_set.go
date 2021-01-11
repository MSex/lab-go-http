package sqlite

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	ProvideConn,

	ProvideUsers,
)
