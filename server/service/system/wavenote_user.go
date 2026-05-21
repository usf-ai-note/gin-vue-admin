
package system

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
    systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

type WavenoteUserService struct {}
// CreateWavenoteUser 创建wavenoteUser表记录
// Author [yourname](https://github.com/yourname)
func (wavenoteUserService *WavenoteUserService) CreateWavenoteUser(ctx context.Context, wavenoteUser *system.WavenoteUser) (err error) {
	err = global.MustGetGlobalDBByDBName("ai_user").Create(wavenoteUser).Error
	return err
}

// DeleteWavenoteUser 删除wavenoteUser表记录
// Author [yourname](https://github.com/yourname)
func (wavenoteUserService *WavenoteUserService)DeleteWavenoteUser(ctx context.Context, id string) (err error) {
	err = global.MustGetGlobalDBByDBName("ai_user").Delete(&system.WavenoteUser{},"id = ?",id).Error
	return err
}

// DeleteWavenoteUserByIds 批量删除wavenoteUser表记录
// Author [yourname](https://github.com/yourname)
func (wavenoteUserService *WavenoteUserService)DeleteWavenoteUserByIds(ctx context.Context, ids []string) (err error) {
	err = global.MustGetGlobalDBByDBName("ai_user").Delete(&[]system.WavenoteUser{},"id in ?",ids).Error
	return err
}

// UpdateWavenoteUser 更新wavenoteUser表记录
// Author [yourname](https://github.com/yourname)
func (wavenoteUserService *WavenoteUserService)UpdateWavenoteUser(ctx context.Context, wavenoteUser system.WavenoteUser) (err error) {
	err = global.MustGetGlobalDBByDBName("ai_user").Model(&system.WavenoteUser{}).Where("id = ?",wavenoteUser.Id).Updates(&wavenoteUser).Error
	return err
}

// GetWavenoteUser 根据id获取wavenoteUser表记录
// Author [yourname](https://github.com/yourname)
func (wavenoteUserService *WavenoteUserService)GetWavenoteUser(ctx context.Context, id string) (wavenoteUser system.WavenoteUser, err error) {
	err = global.MustGetGlobalDBByDBName("ai_user").Where("id = ?", id).First(&wavenoteUser).Error
	return
}
// GetWavenoteUserInfoList 分页获取wavenoteUser表记录
// Author [yourname](https://github.com/yourname)
func (wavenoteUserService *WavenoteUserService)GetWavenoteUserInfoList(ctx context.Context, info systemReq.WavenoteUserSearch) (list []system.WavenoteUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("ai_user").Model(&system.WavenoteUser{})
    var wavenoteUsers []system.WavenoteUser
    // 如果有条件搜索 下方会自动创建搜索语句
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&wavenoteUsers).Error
	return  wavenoteUsers, total, err
}
func (wavenoteUserService *WavenoteUserService)GetWavenoteUserPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
