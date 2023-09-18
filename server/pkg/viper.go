package pkg

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"nebula.xyz/global"
	"nebula.xyz/helper"
)

// 配置Viper
func Viper() *viper.Viper{

	v := viper.New()
	v.SetConfigType(helper.ConfigType)
	v.SetConfigName(helper.ConfigName)
	v.AddConfigPath(helper.ConfigPath)
	err := v.ReadInConfig()
	if err != nil{
		panic(fmt.Errorf("Fatal error config file: %w", err))
	}


	// 监听用户的更改
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		// TODO 当配置文件修改后响应
		fmt.Println("Config file changed:", e.Name)
	})

	if err = v.Unmarshal(&global.CONFING); err != nil {
		fmt.Println(err)
	}

	return v
}