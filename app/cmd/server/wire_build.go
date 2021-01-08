// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/MSex/lab-go-http/app/config"
	"github.com/MSex/lab-go-http/app/data/nodb"
	"github.com/MSex/lab-go-http/app/logger"
	"github.com/MSex/lab-go-http/app/server"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

func injectLogger() (*zap.Logger, error) {
	wire.Build(
		logger.ProviderSet,
		config.ProviderSet,
	)
	return nil, nil
}

func inject() (*httprouter.Router, error) {
	wire.Build(
		logger.ProviderSet,
		config.ProviderSet,
		nodb.ProviderSet,
		server.ProviderSet,
	)
	return nil, nil
}
