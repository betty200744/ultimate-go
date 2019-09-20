package main

import (
	"encoding/json"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func main() {
	now := time.Now()
	cfg := zap.Config{}
	configTemp := `{
		  "level": "DEBUG",
		  "encoding": "console",
		  "outputPaths": ["stdout"],
		  "errorOutputPaths": ["stdout"]
		  }`
	cfg.EncoderConfig = zap.NewProductionEncoderConfig()
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // 添加颜色
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder        // 时间格式
	if err := json.Unmarshal([]byte(configTemp), &cfg); err != nil {
		panic(err)
	}

	logger, _ := cfg.Build()

	logger.Info("this is info",
		zap.Int("a", 1),
		zap.String("b", "b"),
		zap.Any("c", []int{1, 3}),
		zap.Duration("duration", time.Now().Sub(now)))
	logger.Warn("this is warn")
	logger.Debug("this is debug")
	logger.Panic("this is panic")
}
