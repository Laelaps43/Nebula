package internal

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"nebula.xyz/global"
)

type _zap struct{}

var Zap = new(_zap)

// GetZapCores 获取配置Zap的Core
func (z *_zap) GetZapCores() []zapcore.Core {
	cores := make([]zapcore.Core, 0, 4)
	for level := global.CONFIG.ZAP.TransLevel(); level <= zapcore.ErrorLevel; level++ {
		cores = append(cores, z.GetEncoderCore(level, z.GetLevelPriority(level)))
	}
	return cores
}

// GetEncoderCore 获取指定Encoder的zapcore
func (z *_zap) GetEncoderCore(level zapcore.Level, levelFunc zap.LevelEnablerFunc) zapcore.Core {
	writer := FileRotatelogs.GetWriterSyncer(level.String())

	return zapcore.NewCore(z.GetEncoder(), writer, level)
}

// GetEncoder 返回指定的编码器
func (z *_zap) GetEncoder() zapcore.Encoder {
	if global.CONFIG.ZAP.Format == "json" {
		return zapcore.NewJSONEncoder(z.GetEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(z.GetEncoderConfig())
}

// CustomTimeEncoder 自定义日志时间输出格式
func (z *_zap) CustomTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(t.Format("2006/01/02 - 15:04:05.000"))
}

// GetEncoderConfig 自定义配置
func (z *_zap) GetEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey: "message",
		LevelKey:   "level",
		TimeKey:    "time",
		NameKey:    "logger",
		CallerKey:  "caller",
		// StacktraceKey:  global.GVA_CONFIG.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseColorLevelEncoder, // 小写带颜色编码器
		EncodeTime:     z.CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
}

// GetLevelPriority 日志级别
func (z *_zap) GetLevelPriority(level zapcore.Level) zap.LevelEnablerFunc {
	switch level {
	case zapcore.DebugLevel:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	case zapcore.InfoLevel:
		return func(level zapcore.Level) bool { // 日志级别
			return level == zap.InfoLevel
		}
	case zapcore.WarnLevel:
		return func(level zapcore.Level) bool { // 警告级别
			return level == zap.WarnLevel
		}
	case zapcore.ErrorLevel:
		return func(level zapcore.Level) bool { // 错误级别
			return level == zap.ErrorLevel
		}
	default:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	}
}
