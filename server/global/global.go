package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"nebula.xyz/config"
)

type SipConfig interface{}

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

	// SingleFlight 用来处理并发控制
	SingleFlight = &singleflight.Group{}

	MediaServer Media

	Info = &SystemInfo{
		CPUList:  make([]string, 60, 60),
		MemList:  make([]string, 60, 60),
		DiskList: make([]string, 60, 60),
		UpList:   make([]string, 60, 60),
		DownList: make([]string, 60, 60),
		TimeList: make([]string, 60, 60),
	}
)
