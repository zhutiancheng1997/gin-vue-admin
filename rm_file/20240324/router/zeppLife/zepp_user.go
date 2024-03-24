package zeppLife

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ZeppUserRouter struct {
}

// InitZeppUserRouter 初始化 zeppUser 路由信息
func (s *ZeppUserRouter) InitZeppUserRouter(Router *gin.RouterGroup) {
	zeppUserRouter := Router.Group("zeppUser").Use(middleware.OperationRecord())
	zeppUserRouterWithoutRecord := Router.Group("zeppUser")
	var zeppUserApi = v1.ApiGroupApp.ZeppLifeApiGroup.ZeppUserApi
	{
		zeppUserRouter.POST("createZeppUser", zeppUserApi.CreateZeppUser)   // 新建zeppUser
		zeppUserRouter.DELETE("deleteZeppUser", zeppUserApi.DeleteZeppUser) // 删除zeppUser
		zeppUserRouter.DELETE("deleteZeppUserByIds", zeppUserApi.DeleteZeppUserByIds) // 批量删除zeppUser
		zeppUserRouter.PUT("updateZeppUser", zeppUserApi.UpdateZeppUser)    // 更新zeppUser
	}
	{
		zeppUserRouterWithoutRecord.GET("findZeppUser", zeppUserApi.FindZeppUser)        // 根据ID获取zeppUser
		zeppUserRouterWithoutRecord.GET("getZeppUserList", zeppUserApi.GetZeppUserList)  // 获取zeppUser列表
	}
}
