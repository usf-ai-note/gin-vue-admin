package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type WavenoteUserSearch struct {
	Id           *int32  `json:"id" form:"id"`
	Email        *string `json:"email" form:"email"`
	ThirdAccount *string `json:"thirdAccount" form:"thirdAccount"`
	ChannelUid   *string `json:"channelUid" form:"channelUid"`
	Origin       *int32  `json:"origin" form:"origin"`
	SyncStatus   *int32  `json:"syncStatus" form:"syncStatus"`
	request.PageInfo
}
