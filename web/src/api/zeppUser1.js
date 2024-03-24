import service from '@/utils/request'

// @Tags ZeppUser1
// @Summary 创建zeppUser1表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZeppUser1 true "创建zeppUser1表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /zeppUser1/createZeppUser1 [post]
export const createZeppUser1 = (data) => {
  return service({
    url: '/zeppUser1/createZeppUser1',
    method: 'post',
    data
  })
}

// @Tags ZeppUser1
// @Summary 删除zeppUser1表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZeppUser1 true "删除zeppUser1表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /zeppUser1/deleteZeppUser1 [delete]
export const deleteZeppUser1 = (params) => {
  return service({
    url: '/zeppUser1/deleteZeppUser1',
    method: 'delete',
    params
  })
}

// @Tags ZeppUser1
// @Summary 批量删除zeppUser1表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除zeppUser1表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /zeppUser1/deleteZeppUser1 [delete]
export const deleteZeppUser1ByIds = (params) => {
  return service({
    url: '/zeppUser1/deleteZeppUser1ByIds',
    method: 'delete',
    params
  })
}

// @Tags ZeppUser1
// @Summary 更新zeppUser1表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZeppUser1 true "更新zeppUser1表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /zeppUser1/updateZeppUser1 [put]
export const updateZeppUser1 = (data) => {
  return service({
    url: '/zeppUser1/updateZeppUser1',
    method: 'put',
    data
  })
}

// @Tags ZeppUser1
// @Summary 用id查询zeppUser1表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ZeppUser1 true "用id查询zeppUser1表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /zeppUser1/findZeppUser1 [get]
export const findZeppUser1 = (params) => {
  return service({
    url: '/zeppUser1/findZeppUser1',
    method: 'get',
    params
  })
}

// @Tags ZeppUser1
// @Summary 分页获取zeppUser1表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取zeppUser1表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /zeppUser1/getZeppUser1List [get]
export const getZeppUser1List = (params) => {
  return service({
    url: '/zeppUser1/getZeppUser1List',
    method: 'get',
    params
  })
}
