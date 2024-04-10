package web

import (
	"github.com/gin-gonic/gin"
	"nebula.xyz/api/v1/web"
)

type ChannelRouter struct{}

// InitChannelRoute 初始化通道路由
func (c *ChannelRouter) InitChannelRoute(group *gin.RouterGroup) {
	channelRouterGroup := group.Group("channel")
	channelApi := web.WebApiAll.ChannelApi
	{
		channelRouterGroup.POST("list", channelApi.GetAllChannels)
		channelRouterGroup.GET("create/generate", channelApi.GenerateChannel)
		channelRouterGroup.POST("create/create", channelApi.CreateChannel)
		channelRouterGroup.POST("update", channelApi.UpdateChannel)
		channelRouterGroup.GET("delete/:channelId", channelApi.DeleteChannel)
	}
}
