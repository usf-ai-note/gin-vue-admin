package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AudioPptSearch struct {
	request.PageInfo
	Uid     string `json:"uid" form:"uid"`
	AudioId string `json:"audioId" form:"audioId"`
	Status  string `json:"status" form:"status"`
}
