package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"nebula.xyz/config"
)

var (

	// CONFIG 系统配置
	CONFIG config.NEBULA

	// Logger 日志
	Logger *zap.Logger

	// VP viper变量
	VP *viper.Viper

	// DB gorm变量

	DB *gorm.DB

	// CACHE 缓存，可能是Redis缓存，也有可能是本地缓存
	CACHE Cache
)
