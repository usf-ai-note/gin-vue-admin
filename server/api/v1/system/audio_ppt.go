package system

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/system"
    systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type AudioPptApi struct {}



// CreateAudioPpt 创建audioPpt表
// @Tags AudioPpt
// @Summary 创建audioPpt表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.AudioPpt true "创建audioPpt表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /audioPpt/createAudioPpt [post]
func (audioPptApi *AudioPptApi) CreateAudioPpt(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var audioPpt system.AudioPpt
	err := c.ShouldBindJSON(&audioPpt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = audioPptService.CreateAudioPpt(ctx,&audioPpt)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteAudioPpt 删除audioPpt表
// @Tags AudioPpt
// @Summary 删除audioPpt表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.AudioPpt true "删除audioPpt表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /audioPpt/deleteAudioPpt [delete]
func (audioPptApi *AudioPptApi) DeleteAudioPpt(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	err := audioPptService.DeleteAudioPpt(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteAudioPptByIds 批量删除audioPpt表
// @Tags AudioPpt
// @Summary 批量删除audioPpt表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /audioPpt/deleteAudioPptByIds [delete]
func (audioPptApi *AudioPptApi) DeleteAudioPptByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := audioPptService.DeleteAudioPptByIds(ctx,ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateAudioPpt 更新audioPpt表
// @Tags AudioPpt
// @Summary 更新audioPpt表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.AudioPpt true "更新audioPpt表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /audioPpt/updateAudioPpt [put]
func (audioPptApi *AudioPptApi) UpdateAudioPpt(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var audioPpt system.AudioPpt
	err := c.ShouldBindJSON(&audioPpt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = audioPptService.UpdateAudioPpt(ctx,audioPpt)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindAudioPpt 用id查询audioPpt表
// @Tags AudioPpt
// @Summary 用id查询audioPpt表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询audioPpt表"
// @Success 200 {object} response.Response{data=system.AudioPpt,msg=string} "查询成功"
// @Router /audioPpt/findAudioPpt [get]
func (audioPptApi *AudioPptApi) FindAudioPpt(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	reaudioPpt, err := audioPptService.GetAudioPpt(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reaudioPpt, c)
}
// GetAudioPptList 分页获取audioPpt表列表
// @Tags AudioPpt
// @Summary 分页获取audioPpt表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.AudioPptSearch true "分页获取audioPpt表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /audioPpt/getAudioPptList [get]
func (audioPptApi *AudioPptApi) GetAudioPptList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo systemReq.AudioPptSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := audioPptService.GetAudioPptInfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetAudioPptPublic 不需要鉴权的audioPpt表接口
// @Tags AudioPpt
// @Summary 不需要鉴权的audioPpt表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /audioPpt/getAudioPptPublic [get]
func (audioPptApi *AudioPptApi) GetAudioPptPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    audioPptService.GetAudioPptPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的audioPpt表接口信息",
    }, "获取成功", c)
}
