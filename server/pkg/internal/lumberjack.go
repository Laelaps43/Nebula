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
func (r fileRotatelogs) GetWriterSyncer(level string)(zapcore.WriteSyncer){
	fileWriter := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s",global.CONFING.ZAP.Directory,helper.LogName),
		MaxSize:    global.CONFING.ZAP.MaxSize, // megabytes
		MaxBackups: global.CONFING.ZAP.MaxBackups,
		MaxAge:     global.CONFING.ZAP.MaxAge, //days
		Compress:   global.CONFING.ZAP.Compress, // disabled by default
	}
	if global.CONFING.ZAP.LoginConsole{
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter))
	}
	return zapcore.AddSync(fileWriter)
}