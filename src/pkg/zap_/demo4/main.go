package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

func main() {
	logger, err := getLogger("info1.log", "error1.log")
	if err != nil {
		log.Fatal(err)
	}

	logger.Debug("i am debug", zap.String("key", "debug"))
	logger.Info("i am info", zap.String("key", "info"))
	logger.Error("i am error", zap.String("key", "error"))
}

func getLogger(infoPath, errorPath string) (*zap.Logger, error) {
	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})

	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev < zap.ErrorLevel && lev >= zap.DebugLevel
	})

	prodEncoder := zap.NewProductionEncoderConfig()
	prodEncoder.EncodeTime = zapcore.ISO8601TimeEncoder

	lowWriteSyncer, lowClose, err := zap.Open(infoPath)
	if err != nil {
		lowClose()
		return nil, err
	}

	highWriteSyncer, highClose, err := zap.Open(errorPath)
	if err != nil {
		highClose()
		return nil, err
	}

	highCore := zapcore.NewCore(zapcore.NewJSONEncoder(prodEncoder), highWriteSyncer, highPriority)
	lowCore := zapcore.NewCore(zapcore.NewJSONEncoder(prodEncoder), lowWriteSyncer, lowPriority)

	return zap.New(zapcore.NewTee(highCore, lowCore), zap.AddCaller()), nil
}
