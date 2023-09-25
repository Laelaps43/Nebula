package internal

import (
	"fmt"
	"os"

	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"nebula.xyz/global"
	"nebula.xyz/helper"
)

var FileRotatelogs = new(fileRotatelogs)

type fileRotatelogs struct{}

// 使用lumberjack进行日志分割
func (r fileRotatelogs) GetWriterSyncer(level string) zapcore.WriteSyncer {
	fileWriter := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", global.CONFIG.ZAP.Directory, helper.LogName),
		MaxSize:    global.CONFIG.ZAP.MaxSize, // megabytes
		MaxBackups: global.CONFIG.ZAP.MaxBackups,
		MaxAge:     global.CONFIG.ZAP.MaxAge,   //days
		Compress:   global.CONFIG.ZAP.Compress, // disabled by default
	}
	if global.CONFIG.ZAP.LoginConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter))
	}
	return zapcore.AddSync(fileWriter)
}
