package service

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"nebula.xyz/global"
	"sync"
)

type CasbinService struct{}

var (
	syncedCachedEnforcer *casbin.SyncedCachedEnforcer
	once                 sync.Once
)

func (c *CasbinService) Casbin() *casbin.SyncedCachedEnforcer {
	once.Do(func() {
		global.Logger.Info("开始注册Casbin")
		adapter, err := gormadapter.NewAdapterByDB(global.DB)
		if err != nil {
			global.Logger.Error("Casbin 适配数据库设备", zap.Error(err))
			return
		}
		text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
		modelFromString, err := model.NewModelFromString(text)
		if err != nil {
			global.Logger.Error("casbin加载字符串模型失败", zap.Error(err))
			return
		}
		global.Logger.Info("加载字符串模型成功")
		syncedCachedEnforcer, err = casbin.NewSyncedCachedEnforcer(modelFromString, adapter)
		if err != nil {
			global.Logger.Error("创建Casbin错误", zap.Error(err))
		}
		global.Logger.Info("开始加载政策模型.....")
		syncedCachedEnforcer.SetExpireTime(60 * 60)
		err = syncedCachedEnforcer.LoadPolicy()
		if err != nil {
			global.Logger.Error("casbin加载Policy失败", zap.Error(err))
			return
		}
		global.Logger.Info("加载政策模型成功")
	})
	return syncedCachedEnforcer
}
