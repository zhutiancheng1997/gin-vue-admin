package zeppLife

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/zeppLife"
    zeppLifeReq "github.com/flipped-aurora/gin-vue-admin/server/model/zeppLife/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type Test1UserApi struct {
}

var TestUserService = service.ServiceGroupApp.ZeppLifeServiceGroup.Test1UserService


// CreateTest1User 创建Test1User
// @Tags Test1User
// @Summary 创建Test1User
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body zeppLife.Test1User true "创建Test1User"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /TestUser/createTest1User [post]
func (TestUserApi *Test1UserApi) CreateTest1User(c *gin.Context) {
	var TestUser zeppLife.Test1User
	err := c.ShouldBindJSON(&TestUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := TestUserService.CreateTest1User(&TestUser); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTest1User 删除Test1User
// @Tags Test1User
// @Summary 删除Test1User
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body zeppLife.Test1User true "删除Test1User"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TestUser/deleteTest1User [delete]
func (TestUserApi *Test1UserApi) DeleteTest1User(c *gin.Context) {
	ID := c.Query("ID")
	if err := TestUserService.DeleteTest1User(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTest1UserByIds 批量删除Test1User
// @Tags Test1User
// @Summary 批量删除Test1User
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /TestUser/deleteTest1UserByIds [delete]
func (TestUserApi *Test1UserApi) DeleteTest1UserByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := TestUserService.DeleteTest1UserByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTest1User 更新Test1User
// @Tags Test1User
// @Summary 更新Test1User
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body zeppLife.Test1User true "更新Test1User"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /TestUser/updateTest1User [put]
func (TestUserApi *Test1UserApi) UpdateTest1User(c *gin.Context) {
	var TestUser zeppLife.Test1User
	err := c.ShouldBindJSON(&TestUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := TestUserService.UpdateTest1User(TestUser); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTest1User 用id查询Test1User
// @Tags Test1User
// @Summary 用id查询Test1User
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query zeppLife.Test1User true "用id查询Test1User"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /TestUser/findTest1User [get]
func (TestUserApi *Test1UserApi) FindTest1User(c *gin.Context) {
	ID := c.Query("ID")
	if reTestUser, err := TestUserService.GetTest1User(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reTestUser": reTestUser}, c)
	}
}

// GetTest1UserList 分页获取Test1User列表
// @Tags Test1User
// @Summary 分页获取Test1User列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query zeppLifeReq.Test1UserSearch true "分页获取Test1User列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TestUser/getTest1UserList [get]
func (TestUserApi *Test1UserApi) GetTest1UserList(c *gin.Context) {
	var pageInfo zeppLifeReq.Test1UserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := TestUserService.GetTest1UserInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
