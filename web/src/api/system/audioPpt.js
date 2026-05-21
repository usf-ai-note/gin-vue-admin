import service from '@/utils/request'
// @Tags AudioPpt
// @Summary 创建audioPpt表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AudioPpt true "创建audioPpt表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /audioPpt/createAudioPpt [post]
export const createAudioPpt = (data) => {
  return service({
    url: '/audioPpt/createAudioPpt',
    method: 'post',
    data
  })
}

// @Tags AudioPpt
// @Summary 删除audioPpt表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AudioPpt true "删除audioPpt表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /audioPpt/deleteAudioPpt [delete]
export const deleteAudioPpt = (params) => {
  return service({
    url: '/audioPpt/deleteAudioPpt',
    method: 'delete',
    params
  })
}

// @Tags AudioPpt
// @Summary 批量删除audioPpt表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除audioPpt表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /audioPpt/deleteAudioPpt [delete]
export const deleteAudioPptByIds = (params) => {
  return service({
    url: '/audioPpt/deleteAudioPptByIds',
    method: 'delete',
    params
  })
}

// @Tags AudioPpt
// @Summary 更新audioPpt表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AudioPpt true "更新audioPpt表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /audioPpt/updateAudioPpt [put]
export const updateAudioPpt = (data) => {
  return service({
    url: '/audioPpt/updateAudioPpt',
    method: 'put',
    data
  })
}

// @Tags AudioPpt
// @Summary 用id查询audioPpt表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.AudioPpt true "用id查询audioPpt表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /audioPpt/findAudioPpt [get]
export const findAudioPpt = (params) => {
  return service({
    url: '/audioPpt/findAudioPpt',
    method: 'get',
    params
  })
}

// @Tags AudioPpt
// @Summary 分页获取audioPpt表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取audioPpt表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /audioPpt/getAudioPptList [get]
export const getAudioPptList = (params) => {
  return service({
    url: '/audioPpt/getAudioPptList',
    method: 'get',
    params
  })
}

// @Tags AudioPpt
// @Summary 不需要鉴权的audioPpt表接口
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.AudioPptSearch true "分页获取audioPpt表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /audioPpt/getAudioPptPublic [get]
export const getAudioPptPublic = () => {
  return service({
    url: '/audioPpt/getAudioPptPublic',
    method: 'get',
  })
}
