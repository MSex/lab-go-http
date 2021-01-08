package config

import (
	"io/ioutil"
	"strings"

	"github.com/MSex/lab-go-http/app/logger"
	"go.uber.org/zap/zapcore"
)

const appName = logger.AppName("lab-go-http")

// TODO replace this
const logLevel = zapcore.DebugLevel
const devLogging = logger.DevelopmentLogging(true)

func ProvideAppName() logger.AppName {
	return appName
}

func ProvideBuildNumber() (logger.BuildNumber, error) {
	buf, err := ioutil.ReadFile("res/build_number")
	if err != nil {
		return "", err
	}
	buildNumber := strings.Replace(string(buf), "\n", "", -1)
	return logger.BuildNumber(buildNumber), nil
}


func ProvideLogLevel() zapcore.Level {
	return logLevel
}

func ProvideDevLogging() logger.DevelopmentLogging {
	return devLogging
}
