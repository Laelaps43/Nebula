package pkg

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"nebula.xyz/global"
	"nebula.xyz/pkg/internal"
	"nebula.xyz/utils"
)

// zap配置
func Zap() (logger *zap.Logger) {
	var directory = global.CONFIG.ZAP.Directory
	if ok, _ := utils.PathExists(directory); !ok {
		// 判断文件夹下是否存在指定的文件夹
		fmt.Printf("指定的文件夹%s不存在，创建文件夹中...\n", directory)
		os.Mkdir(directory, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	// TODO 是否增加Liner Number
	return logger
}
