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

type XmUserApi struct {
}

var xmUserService = service.ServiceGroupApp.ZeppLifeServiceGroup.XmUserService


// CreateXmUser 创建xmUser
// @Tags XmUser
// @Summary 创建xmUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body zeppLife.XmUser true "创建xmUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /xmUser/createXmUser [post]
func (xmUserApi *XmUserApi) CreateXmUser(c *gin.Context) {
	var xmUser zeppLife.XmUser
	err := c.ShouldBindJSON(&xmUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := xmUserService.CreateXmUser(&xmUser); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteXmUser 删除xmUser
// @Tags XmUser
// @Summary 删除xmUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body zeppLife.XmUser true "删除xmUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /xmUser/deleteXmUser [delete]
func (xmUserApi *XmUserApi) DeleteXmUser(c *gin.Context) {
	ID := c.Query("ID")
	if err := xmUserService.DeleteXmUser(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteXmUserByIds 批量删除xmUser
// @Tags XmUser
// @Summary 批量删除xmUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /xmUser/deleteXmUserByIds [delete]
func (xmUserApi *XmUserApi) DeleteXmUserByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := xmUserService.DeleteXmUserByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateXmUser 更新xmUser
// @Tags XmUser
// @Summary 更新xmUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body zeppLife.XmUser true "更新xmUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /xmUser/updateXmUser [put]
func (xmUserApi *XmUserApi) UpdateXmUser(c *gin.Context) {
	var xmUser zeppLife.XmUser
	err := c.ShouldBindJSON(&xmUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := xmUserService.UpdateXmUser(xmUser); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindXmUser 用id查询xmUser
// @Tags XmUser
// @Summary 用id查询xmUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query zeppLife.XmUser true "用id查询xmUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /xmUser/findXmUser [get]
func (xmUserApi *XmUserApi) FindXmUser(c *gin.Context) {
	ID := c.Query("ID")
	if rexmUser, err := xmUserService.GetXmUser(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rexmUser": rexmUser}, c)
	}
}

// GetXmUserList 分页获取xmUser列表
// @Tags XmUser
// @Summary 分页获取xmUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query zeppLifeReq.XmUserSearch true "分页获取xmUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /xmUser/getXmUserList [get]
func (xmUserApi *XmUserApi) GetXmUserList(c *gin.Context) {
	var pageInfo zeppLifeReq.XmUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := xmUserService.GetXmUserInfoList(pageInfo); err != nil {
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
