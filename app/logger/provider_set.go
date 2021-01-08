package logger

import (
	"github.com/google/wire"
)

// ProviderSet represents the injection module
var ProviderSet = wire.NewSet(
	ProvideLoggerConfig,
	ProvideLogger,
)
