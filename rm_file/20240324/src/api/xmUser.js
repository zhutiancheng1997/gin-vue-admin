import service from '@/utils/request'

// @Tags XmUser
// @Summary 创建xmUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.XmUser true "创建xmUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /xmUser/createXmUser [post]
export const createXmUser = (data) => {
  return service({
    url: '/xmUser/createXmUser',
    method: 'post',
    data
  })
}

// @Tags XmUser
// @Summary 删除xmUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.XmUser true "删除xmUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /xmUser/deleteXmUser [delete]
export const deleteXmUser = (params) => {
  return service({
    url: '/xmUser/deleteXmUser',
    method: 'delete',
    params
  })
}

// @Tags XmUser
// @Summary 批量删除xmUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除xmUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /xmUser/deleteXmUser [delete]
export const deleteXmUserByIds = (params) => {
  return service({
    url: '/xmUser/deleteXmUserByIds',
    method: 'delete',
    params
  })
}

// @Tags XmUser
// @Summary 更新xmUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.XmUser true "更新xmUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /xmUser/updateXmUser [put]
export const updateXmUser = (data) => {
  return service({
    url: '/xmUser/updateXmUser',
    method: 'put',
    data
  })
}

// @Tags XmUser
// @Summary 用id查询xmUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.XmUser true "用id查询xmUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /xmUser/findXmUser [get]
export const findXmUser = (params) => {
  return service({
    url: '/xmUser/findXmUser',
    method: 'get',
    params
  })
}

// @Tags XmUser
// @Summary 分页获取xmUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取xmUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /xmUser/getXmUserList [get]
export const getXmUserList = (params) => {
  return service({
    url: '/xmUser/getXmUserList',
    method: 'get',
    params
  })
}
