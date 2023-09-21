package config
// zap配置信息

import (
	"strings"

	"go.uber.org/zap/zapcore"
)

type ZAP struct {
	Directory    string `yaml:"directory"`
	Level        string `yaml:"level"`
	MaxSize      int    `yaml:"maxSize"`
	MaxBackups   int    `yaml:"maxBackups"`
	MaxAge       int    `yaml:"maxAge"`
	Compress     bool   `yaml:"compress"`
	LoginConsole bool   `yaml:"loginConsole"`
	Format       string `yaml:"format"`
}

// 根据指定的类型返回对应zapcore的等级
func (zap *ZAP) TransLevel() zapcore.Level {
	zap.Level = strings.ToLower(zap.Level)
	switch zap.Level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.DebugLevel
	}
}
