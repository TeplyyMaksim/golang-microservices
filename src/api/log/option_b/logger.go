package option_b

import (
	"github.com/uber-go/zap"
	"github.com/uber-go/zap/zapcore"
)

var (
	Log *zap.Logger
)

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{"stdout", "/tmp/logs"},
		Encoding: "json",
		Level: zap.NewAtomicLevelAt(zap.InfoLevel),
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:       "msg",
			LevelKey:         "level",
			TimeKey:          "time",
			EncodeTime: 		zapcore.ISO8601TimeEncoder,
			EncodeLevel: zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	var err error
	Log, err = logConfig.Build()

	if err != nil {
		panic(err)
	}
}

func Field(key string, value interface{}) zap.Field {
	return zap.Array(key, value)
}

func Debug(msg string, tags ...zap.Field) {
	Log.Debug(msg, tags...)
	Log.Sync()
}

func Info(msg string, tags ...zap.Field) {
	Log.Info(msg, tags...)
	Log.Sync()
}

func Error(msg string, tags ...zap.Field) {
	Log.Error(msg, tags...)
	Log.Sync()
}