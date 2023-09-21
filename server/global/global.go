package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"nebula.xyz/config"
)


var (
	
	// 系统配置
	CONFING 	config.NEBULA;


	// 日志
	Logger		*zap.Logger

	// viper变量
	VP			*viper.Viper

	// gorm变量
	DB			*gorm.DB
)


