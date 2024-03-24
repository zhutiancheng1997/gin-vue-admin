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

type ZeppUserApi struct {
}

var zeppUserService = service.ServiceGroupApp.ZeppLifeServiceGroup.ZeppUserService


// CreateZeppUser 创建zeppUser
// @Tags ZeppUser
// @Summary 创建zeppUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body zeppLife.ZeppUser true "创建zeppUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /zeppUser/createZeppUser [post]
func (zeppUserApi *ZeppUserApi) CreateZeppUser(c *gin.Context) {
	var zeppUser zeppLife.ZeppUser
	err := c.ShouldBindJSON(&zeppUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := zeppUserService.CreateZeppUser(&zeppUser); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteZeppUser 删除zeppUser
// @Tags ZeppUser
// @Summary 删除zeppUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body zeppLife.ZeppUser true "删除zeppUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /zeppUser/deleteZeppUser [delete]
func (zeppUserApi *ZeppUserApi) DeleteZeppUser(c *gin.Context) {
	ID := c.Query("ID")
	if err := zeppUserService.DeleteZeppUser(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteZeppUserByIds 批量删除zeppUser
// @Tags ZeppUser
// @Summary 批量删除zeppUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /zeppUser/deleteZeppUserByIds [delete]
func (zeppUserApi *ZeppUserApi) DeleteZeppUserByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := zeppUserService.DeleteZeppUserByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateZeppUser 更新zeppUser
// @Tags ZeppUser
// @Summary 更新zeppUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body zeppLife.ZeppUser true "更新zeppUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /zeppUser/updateZeppUser [put]
func (zeppUserApi *ZeppUserApi) UpdateZeppUser(c *gin.Context) {
	var zeppUser zeppLife.ZeppUser
	err := c.ShouldBindJSON(&zeppUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := zeppUserService.UpdateZeppUser(zeppUser); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindZeppUser 用id查询zeppUser
// @Tags ZeppUser
// @Summary 用id查询zeppUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query zeppLife.ZeppUser true "用id查询zeppUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /zeppUser/findZeppUser [get]
func (zeppUserApi *ZeppUserApi) FindZeppUser(c *gin.Context) {
	ID := c.Query("ID")
	if rezeppUser, err := zeppUserService.GetZeppUser(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rezeppUser": rezeppUser}, c)
	}
}

// GetZeppUserList 分页获取zeppUser列表
// @Tags ZeppUser
// @Summary 分页获取zeppUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query zeppLifeReq.ZeppUserSearch true "分页获取zeppUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /zeppUser/getZeppUserList [get]
func (zeppUserApi *ZeppUserApi) GetZeppUserList(c *gin.Context) {
	var pageInfo zeppLifeReq.ZeppUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := zeppUserService.GetZeppUserInfoList(pageInfo); err != nil {
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
