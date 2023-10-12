/*
 * @Date: 2023-10-12 01:15:03
 * @LastEditors: HeXu
 * @LastEditTime: 2023-10-12 01:40:16
 * @FilePath: /crawler/log/default.go
 */
package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)


func DefaultEncodingConfig() zapcore.EncoderConfig {
	var encoderConfig = zap.NewProductionEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	return encoderConfig
}

func DefaultOption() []zap.Option {
	var stackTraceLevel zap.LevelEnablerFunc = func(level zapcore.Level) bool{
		return level >= zapcore.DPanicLevel
	}
	return []zap.Option{
		zap.AddCaller(),
		zap.AddStacktrace(stackTraceLevel),
	}
}

func DefaultEncoder() zapcore.Encoder{
	return zapcore.NewConsoleEncoder(DefaultEncodingConfig())
}

func DefaultLumberjackLogger() *lumberjack.Logger{
	return &lumberjack.Logger{
		MaxSize:200,
		LocalTime: true,
		Compress: true,
	}
}