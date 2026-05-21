package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WavenoteUserRouter struct {}

// InitWavenoteUserRouter 初始化 wavenoteUser表 路由信息
func (s *WavenoteUserRouter) InitWavenoteUserRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	wavenoteUserRouter := Router.Group("wavenoteUser").Use(middleware.OperationRecord())
	wavenoteUserRouterWithoutRecord := Router.Group("wavenoteUser")
	wavenoteUserRouterWithoutAuth := PublicRouter.Group("wavenoteUser")
	{
		wavenoteUserRouter.POST("createWavenoteUser", wavenoteUserApi.CreateWavenoteUser)   // 新建wavenoteUser表
		wavenoteUserRouter.DELETE("deleteWavenoteUser", wavenoteUserApi.DeleteWavenoteUser) // 删除wavenoteUser表
		wavenoteUserRouter.DELETE("deleteWavenoteUserByIds", wavenoteUserApi.DeleteWavenoteUserByIds) // 批量删除wavenoteUser表
		wavenoteUserRouter.PUT("updateWavenoteUser", wavenoteUserApi.UpdateWavenoteUser)    // 更新wavenoteUser表
	}
	{
		wavenoteUserRouterWithoutRecord.GET("findWavenoteUser", wavenoteUserApi.FindWavenoteUser)        // 根据ID获取wavenoteUser表
		wavenoteUserRouterWithoutRecord.GET("getWavenoteUserList", wavenoteUserApi.GetWavenoteUserList)  // 获取wavenoteUser表列表
	}
	{
	    wavenoteUserRouterWithoutAuth.GET("getWavenoteUserPublic", wavenoteUserApi.GetWavenoteUserPublic)  // wavenoteUser表开放接口
	}
}
