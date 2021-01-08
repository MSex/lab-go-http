package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// AppName is the app name
type AppName string

// BuildNumber is the build number
type BuildNumber string

// DevelopmentLogging logger in dev mode
type DevelopmentLogging bool

type AppInfo struct {
	Name        AppName
	BuildNumber BuildNumber
}

// ProvideLoggerConfig builds the logger config
func ProvideLoggerConfig(
	dev DevelopmentLogging,
	level zapcore.Level,
	name AppName,
	build BuildNumber,
) *zap.Config {
	var config = &zap.Config{
		Level:             zap.NewAtomicLevelAt(level),
		Development:       bool(dev),
		DisableCaller:     false,
		DisableStacktrace: true,
		Sampling:          nil,
		Encoding:          "console",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:       "msg",
			LevelKey:         "level",
			TimeKey:          "ts",
			NameKey:          "logger",
			CallerKey:        "caller",
			FunctionKey:      "function",
			StacktraceKey:    "stack",
			LineEnding:       zapcore.DefaultLineEnding,
			EncodeLevel:      zapcore.LowercaseLevelEncoder,
			EncodeTime:       zapcore.ISO8601TimeEncoder,
			EncodeDuration:   zapcore.StringDurationEncoder,
			EncodeCaller:     zapcore.ShortCallerEncoder,
			EncodeName:       zapcore.FullNameEncoder,
			ConsoleSeparator: "\t",
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
		InitialFields: map[string]interface{}{
			"appName":     name,
			"buildNumber": build,
		},
	}
	return config
}

// ProvideLogger builds a logger
func ProvideLogger(config *zap.Config) (*zap.Logger, error) {
	return config.Build()
}
