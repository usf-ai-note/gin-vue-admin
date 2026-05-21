import service from '@/utils/request'
// @Tags WavenoteUser
// @Summary 创建wavenoteUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.WavenoteUser true "创建wavenoteUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wavenoteUser/createWavenoteUser [post]
export const createWavenoteUser = (data) => {
  return service({
    url: '/wavenoteUser/createWavenoteUser',
    method: 'post',
    data
  })
}

// @Tags WavenoteUser
// @Summary 删除wavenoteUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.WavenoteUser true "删除wavenoteUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wavenoteUser/deleteWavenoteUser [delete]
export const deleteWavenoteUser = (params) => {
  return service({
    url: '/wavenoteUser/deleteWavenoteUser',
    method: 'delete',
    params
  })
}

// @Tags WavenoteUser
// @Summary 批量删除wavenoteUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除wavenoteUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wavenoteUser/deleteWavenoteUser [delete]
export const deleteWavenoteUserByIds = (params) => {
  return service({
    url: '/wavenoteUser/deleteWavenoteUserByIds',
    method: 'delete',
    params
  })
}

// @Tags WavenoteUser
// @Summary 更新wavenoteUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.WavenoteUser true "更新wavenoteUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wavenoteUser/updateWavenoteUser [put]
export const updateWavenoteUser = (data) => {
  return service({
    url: '/wavenoteUser/updateWavenoteUser',
    method: 'put',
    data
  })
}

// @Tags WavenoteUser
// @Summary 用id查询wavenoteUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.WavenoteUser true "用id查询wavenoteUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wavenoteUser/findWavenoteUser [get]
export const findWavenoteUser = (params) => {
  return service({
    url: '/wavenoteUser/findWavenoteUser',
    method: 'get',
    params
  })
}

// @Tags WavenoteUser
// @Summary 分页获取wavenoteUser表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取wavenoteUser表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wavenoteUser/getWavenoteUserList [get]
export const getWavenoteUserList = (params) => {
  return service({
    url: '/wavenoteUser/getWavenoteUserList',
    method: 'get',
    params
  })
}

// @Tags WavenoteUser
// @Summary 不需要鉴权的wavenoteUser表接口
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.WavenoteUserSearch true "分页获取wavenoteUser表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /wavenoteUser/getWavenoteUserPublic [get]
export const getWavenoteUserPublic = () => {
  return service({
    url: '/wavenoteUser/getWavenoteUserPublic',
    method: 'get',
  })
}
