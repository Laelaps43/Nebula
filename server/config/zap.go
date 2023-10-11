package config

// zap配置信息

import (
	"strings"

	"go.uber.org/zap/zapcore"
)

type ZAP struct {
	Directory    string `mapstructure:"directory" yaml:"directory"`
	Level        string `mapstructure:"level" yaml:"level"`
	MaxSize      int    `mapstructure:"max-size" yaml:"max-size"`
	MaxBackups   int    `mapstructure:"max-backups" yaml:"max-backups"`
	MaxAge       int    `mapstructure:"max-age" yaml:"max-age"`
	Compress     bool   `mapstructure:"compress" yaml:"compress"`
	LoginConsole bool   `mapstructure:"login-console" yaml:"login-console"`
	Format       string `mapstructure:"format" yaml:"format"`
}

// TransLevel 根据指定的类型返回对应zapcore的等级
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
