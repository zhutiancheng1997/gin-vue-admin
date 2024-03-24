package zeppLife

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type XmUserRouter struct {
}

// InitXmUserRouter 初始化 xmUser 路由信息
func (s *XmUserRouter) InitXmUserRouter(Router *gin.RouterGroup) {
	xmUserRouter := Router.Group("xmUser").Use(middleware.OperationRecord())
	xmUserRouterWithoutRecord := Router.Group("xmUser")
	var xmUserApi = v1.ApiGroupApp.ZeppLifeApiGroup.XmUserApi
	{
		xmUserRouter.POST("createXmUser", xmUserApi.CreateXmUser)   // 新建xmUser
		xmUserRouter.DELETE("deleteXmUser", xmUserApi.DeleteXmUser) // 删除xmUser
		xmUserRouter.DELETE("deleteXmUserByIds", xmUserApi.DeleteXmUserByIds) // 批量删除xmUser
		xmUserRouter.PUT("updateXmUser", xmUserApi.UpdateXmUser)    // 更新xmUser
	}
	{
		xmUserRouterWithoutRecord.GET("findXmUser", xmUserApi.FindXmUser)        // 根据ID获取xmUser
		xmUserRouterWithoutRecord.GET("getXmUserList", xmUserApi.GetXmUserList)  // 获取xmUser列表
	}
}
