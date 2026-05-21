package system

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/system"
    systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type WavenoteUserApi struct {}



// CreateWavenoteUser 创建wavenoteUser表
// @Tags WavenoteUser
// @Summary 创建wavenoteUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.WavenoteUser true "创建wavenoteUser表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /wavenoteUser/createWavenoteUser [post]
func (wavenoteUserApi *WavenoteUserApi) CreateWavenoteUser(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var wavenoteUser system.WavenoteUser
	err := c.ShouldBindJSON(&wavenoteUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = wavenoteUserService.CreateWavenoteUser(ctx,&wavenoteUser)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteWavenoteUser 删除wavenoteUser表
// @Tags WavenoteUser
// @Summary 删除wavenoteUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.WavenoteUser true "删除wavenoteUser表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /wavenoteUser/deleteWavenoteUser [delete]
func (wavenoteUserApi *WavenoteUserApi) DeleteWavenoteUser(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	err := wavenoteUserService.DeleteWavenoteUser(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteWavenoteUserByIds 批量删除wavenoteUser表
// @Tags WavenoteUser
// @Summary 批量删除wavenoteUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /wavenoteUser/deleteWavenoteUserByIds [delete]
func (wavenoteUserApi *WavenoteUserApi) DeleteWavenoteUserByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := wavenoteUserService.DeleteWavenoteUserByIds(ctx,ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateWavenoteUser 更新wavenoteUser表
// @Tags WavenoteUser
// @Summary 更新wavenoteUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.WavenoteUser true "更新wavenoteUser表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /wavenoteUser/updateWavenoteUser [put]
func (wavenoteUserApi *WavenoteUserApi) UpdateWavenoteUser(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var wavenoteUser system.WavenoteUser
	err := c.ShouldBindJSON(&wavenoteUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = wavenoteUserService.UpdateWavenoteUser(ctx,wavenoteUser)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindWavenoteUser 用id查询wavenoteUser表
// @Tags WavenoteUser
// @Summary 用id查询wavenoteUser表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询wavenoteUser表"
// @Success 200 {object} response.Response{data=system.WavenoteUser,msg=string} "查询成功"
// @Router /wavenoteUser/findWavenoteUser [get]
func (wavenoteUserApi *WavenoteUserApi) FindWavenoteUser(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	rewavenoteUser, err := wavenoteUserService.GetWavenoteUser(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(rewavenoteUser, c)
}
// GetWavenoteUserList 分页获取wavenoteUser表列表
// @Tags WavenoteUser
// @Summary 分页获取wavenoteUser表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.WavenoteUserSearch true "分页获取wavenoteUser表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /wavenoteUser/getWavenoteUserList [get]
func (wavenoteUserApi *WavenoteUserApi) GetWavenoteUserList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo systemReq.WavenoteUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := wavenoteUserService.GetWavenoteUserInfoList(ctx,pageInfo)
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

// GetWavenoteUserPublic 不需要鉴权的wavenoteUser表接口
// @Tags WavenoteUser
// @Summary 不需要鉴权的wavenoteUser表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /wavenoteUser/getWavenoteUserPublic [get]
func (wavenoteUserApi *WavenoteUserApi) GetWavenoteUserPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    wavenoteUserService.GetWavenoteUserPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的wavenoteUser表接口信息",
    }, "获取成功", c)
}
