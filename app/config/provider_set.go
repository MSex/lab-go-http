package config

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	ProvideAppName,
	ProvideBuildNumber,
	ProvideDevLogging,
	ProvideLogLevel,
)
