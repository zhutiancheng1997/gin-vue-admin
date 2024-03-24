package zeppLife

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type Test1UserRouter struct {
}

// InitTest1UserRouter 初始化 Test1User 路由信息
func (s *Test1UserRouter) InitTest1UserRouter(Router *gin.RouterGroup) {
	TestUserRouter := Router.Group("TestUser").Use(middleware.OperationRecord())
	TestUserRouterWithoutRecord := Router.Group("TestUser")
	var TestUserApi = v1.ApiGroupApp.ZeppLifeApiGroup.Test1UserApi
	{
		TestUserRouter.POST("createTest1User", TestUserApi.CreateTest1User)   // 新建Test1User
		TestUserRouter.DELETE("deleteTest1User", TestUserApi.DeleteTest1User) // 删除Test1User
		TestUserRouter.DELETE("deleteTest1UserByIds", TestUserApi.DeleteTest1UserByIds) // 批量删除Test1User
		TestUserRouter.PUT("updateTest1User", TestUserApi.UpdateTest1User)    // 更新Test1User
	}
	{
		TestUserRouterWithoutRecord.GET("findTest1User", TestUserApi.FindTest1User)        // 根据ID获取Test1User
		TestUserRouterWithoutRecord.GET("getTest1UserList", TestUserApi.GetTest1UserList)  // 获取Test1User列表
	}
}
