
package system

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
    systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

type AudioPptService struct {}
// CreateAudioPpt 创建audioPpt表记录
// Author [yourname](https://github.com/yourname)
func (audioPptService *AudioPptService) CreateAudioPpt(ctx context.Context, audioPpt *system.AudioPpt) (err error) {
	err = global.MustGetGlobalDBByDBName("ai_user").Create(audioPpt).Error
	return err
}

// DeleteAudioPpt 删除audioPpt表记录
// Author [yourname](https://github.com/yourname)
func (audioPptService *AudioPptService)DeleteAudioPpt(ctx context.Context, id string) (err error) {
	err = global.MustGetGlobalDBByDBName("ai_user").Delete(&system.AudioPpt{},"id = ?",id).Error
	return err
}

// DeleteAudioPptByIds 批量删除audioPpt表记录
// Author [yourname](https://github.com/yourname)
func (audioPptService *AudioPptService)DeleteAudioPptByIds(ctx context.Context, ids []string) (err error) {
	err = global.MustGetGlobalDBByDBName("ai_user").Delete(&[]system.AudioPpt{},"id in ?",ids).Error
	return err
}

// UpdateAudioPpt 更新audioPpt表记录
// Author [yourname](https://github.com/yourname)
func (audioPptService *AudioPptService)UpdateAudioPpt(ctx context.Context, audioPpt system.AudioPpt) (err error) {
	err = global.MustGetGlobalDBByDBName("ai_user").Model(&system.AudioPpt{}).Where("id = ?",audioPpt.Id).Updates(&audioPpt).Error
	return err
}

// GetAudioPpt 根据id获取audioPpt表记录
// Author [yourname](https://github.com/yourname)
func (audioPptService *AudioPptService)GetAudioPpt(ctx context.Context, id string) (audioPpt system.AudioPpt, err error) {
	err = global.MustGetGlobalDBByDBName("ai_user").Where("id = ?", id).First(&audioPpt).Error
	return
}
// GetAudioPptInfoList 分页获取audioPpt表记录
// Author [yourname](https://github.com/yourname)
func (audioPptService *AudioPptService)GetAudioPptInfoList(ctx context.Context, info systemReq.AudioPptSearch) (list []system.AudioPpt, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("ai_user").Model(&system.AudioPpt{})
    var audioPpts []system.AudioPpt
    // 如果有条件搜索 下方会自动创建搜索语句
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&audioPpts).Error
	return  audioPpts, total, err
}
func (audioPptService *AudioPptService)GetAudioPptPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
