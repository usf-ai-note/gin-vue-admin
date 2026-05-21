package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AudioPptRouter struct {}

// InitAudioPptRouter 初始化 audioPpt表 路由信息
func (s *AudioPptRouter) InitAudioPptRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	audioPptRouter := Router.Group("audioPpt").Use(middleware.OperationRecord())
	audioPptRouterWithoutRecord := Router.Group("audioPpt")
	audioPptRouterWithoutAuth := PublicRouter.Group("audioPpt")
	{
		audioPptRouter.POST("createAudioPpt", audioPptApi.CreateAudioPpt)   // 新建audioPpt表
		audioPptRouter.DELETE("deleteAudioPpt", audioPptApi.DeleteAudioPpt) // 删除audioPpt表
		audioPptRouter.DELETE("deleteAudioPptByIds", audioPptApi.DeleteAudioPptByIds) // 批量删除audioPpt表
		audioPptRouter.PUT("updateAudioPpt", audioPptApi.UpdateAudioPpt)    // 更新audioPpt表
	}
	{
		audioPptRouterWithoutRecord.GET("findAudioPpt", audioPptApi.FindAudioPpt)        // 根据ID获取audioPpt表
		audioPptRouterWithoutRecord.GET("getAudioPptList", audioPptApi.GetAudioPptList)  // 获取audioPpt表列表
	}
	{
	    audioPptRouterWithoutAuth.GET("getAudioPptPublic", audioPptApi.GetAudioPptPublic)  // audioPpt表开放接口
	}
}
