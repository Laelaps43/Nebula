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
		channelRouterGroup.GET("", channelApi.GetAllChannels)
		channelRouterGroup.GET(":channelId", channelApi.GetChannelInfoById)
		channelRouterGroup.PUT("", channelApi.UpdateChannelInfo)
		channelRouterGroup.POST("generate", channelApi.GenerateChannel)
	}
}
