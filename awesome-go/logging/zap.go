package main

import (
	"go.uber.org/zap"
	"time"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	now := time.Now()
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", "url",
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", "url")
	sugar.Info("hello")
	logger.Info("this is logger",
		zap.Int("a", 1),
		zap.String("b", "b"),
		zap.Any("c", []int{1, 3}),
		zap.Duration("duration", time.Now().Sub(now)))
}
