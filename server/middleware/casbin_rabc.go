package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"nebula.xyz/global"
	"nebula.xyz/model"
	"nebula.xyz/service"
	"nebula.xyz/utils"
	"strconv"
	"strings"
)

var casbinService = service.SysServiceAll

// CasbinHandler 权限认证
func CasbinHandler() gin.HandlerFunc {

	return func(c *gin.Context) {
		if global.CONFIG.SERVER.Mode == "dev" {
			global.Logger.Info("开始权限认证")
			waitUser, err := utils.GetClaims(c)
			if err != nil {
				model.ErrorWithDetailed(gin.H{}, "没有获取到Token", c)
				c.Abort()
				return
			}
			//请求的路径
			path := c.Request.URL.Path
			obj := strings.TrimPrefix(path, global.CONFIG.SERVER.RouterPrefix)
			// 请求的方法
			act := c.Request.Method

			// 获取角色ID
			sub := strconv.Itoa(int(waitUser.RoleId))
			enforce := casbinService.Casbin()
			global.Logger.Info("casbin获取成功")
			global.Logger.Info(fmt.Sprintf("sub %s, obj %s, act %s", sub, obj, act))
			success, err := enforce.Enforce(sub, obj, act)
			if err != nil {
				global.Logger.Error("鉴权错误", zap.Error(err))
			}
			if !success {
				model.ErrorWithDetailed(gin.H{}, "权限不足", c)
				c.Abort()
				return
			}
			global.Logger.Info("鉴权成功")
		}
		c.Next()
	}
}
