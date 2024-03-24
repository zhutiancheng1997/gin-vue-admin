import service from '@/utils/request'

// @Tags Test1User
// @Summary 创建Test1User
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Test1User true "创建Test1User"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /TestUser/createTest1User [post]
export const createTest1User = (data) => {
  return service({
    url: '/TestUser/createTest1User',
    method: 'post',
    data
  })
}

// @Tags Test1User
// @Summary 删除Test1User
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Test1User true "删除Test1User"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TestUser/deleteTest1User [delete]
export const deleteTest1User = (params) => {
  return service({
    url: '/TestUser/deleteTest1User',
    method: 'delete',
    params
  })
}

// @Tags Test1User
// @Summary 批量删除Test1User
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Test1User"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TestUser/deleteTest1User [delete]
export const deleteTest1UserByIds = (params) => {
  return service({
    url: '/TestUser/deleteTest1UserByIds',
    method: 'delete',
    params
  })
}

// @Tags Test1User
// @Summary 更新Test1User
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Test1User true "更新Test1User"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /TestUser/updateTest1User [put]
export const updateTest1User = (data) => {
  return service({
    url: '/TestUser/updateTest1User',
    method: 'put',
    data
  })
}

// @Tags Test1User
// @Summary 用id查询Test1User
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Test1User true "用id查询Test1User"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /TestUser/findTest1User [get]
export const findTest1User = (params) => {
  return service({
    url: '/TestUser/findTest1User',
    method: 'get',
    params
  })
}

// @Tags Test1User
// @Summary 分页获取Test1User列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Test1User列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TestUser/getTest1UserList [get]
export const getTest1UserList = (params) => {
  return service({
    url: '/TestUser/getTest1UserList',
    method: 'get',
    params
  })
}
