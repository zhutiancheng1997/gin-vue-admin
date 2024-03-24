import service from '@/utils/request'

// @Tags ZeppUser
// @Summary 创建zeppUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZeppUser true "创建zeppUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /zeppUser/createZeppUser [post]
export const createZeppUser = (data) => {
  return service({
    url: '/zeppUser/createZeppUser',
    method: 'post',
    data
  })
}

// @Tags ZeppUser
// @Summary 删除zeppUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZeppUser true "删除zeppUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /zeppUser/deleteZeppUser [delete]
export const deleteZeppUser = (params) => {
  return service({
    url: '/zeppUser/deleteZeppUser',
    method: 'delete',
    params
  })
}

// @Tags ZeppUser
// @Summary 批量删除zeppUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除zeppUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /zeppUser/deleteZeppUser [delete]
export const deleteZeppUserByIds = (params) => {
  return service({
    url: '/zeppUser/deleteZeppUserByIds',
    method: 'delete',
    params
  })
}

// @Tags ZeppUser
// @Summary 更新zeppUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZeppUser true "更新zeppUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /zeppUser/updateZeppUser [put]
export const updateZeppUser = (data) => {
  return service({
    url: '/zeppUser/updateZeppUser',
    method: 'put',
    data
  })
}

// @Tags ZeppUser
// @Summary 用id查询zeppUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ZeppUser true "用id查询zeppUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /zeppUser/findZeppUser [get]
export const findZeppUser = (params) => {
  return service({
    url: '/zeppUser/findZeppUser',
    method: 'get',
    params
  })
}

// @Tags ZeppUser
// @Summary 分页获取zeppUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取zeppUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /zeppUser/getZeppUserList [get]
export const getZeppUserList = (params) => {
  return service({
    url: '/zeppUser/getZeppUserList',
    method: 'get',
    params
  })
}
