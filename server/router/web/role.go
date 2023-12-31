package web

import (
	"github.com/gin-gonic/gin"
	"nebula.xyz/api/v1/web"
)

type RoleRouter struct{}

func (a *RoleRouter) InitRoleRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	roleRouter := Router.Group("/role")
	roleApi := web.WebApiAll.RoleApi
	{
		roleRouter.POST("/create", roleApi.CreateRole)
		roleRouter.POST("/list", roleApi.GetRolePagination)
		roleRouter.POST("/update", roleApi.UpdateRole)
		roleRouter.GET("/delete/:roleId", roleApi.DeleteRole)
	}
	return roleRouter
}
