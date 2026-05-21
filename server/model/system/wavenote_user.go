// 自动生成模板WavenoteUser
package system

import (
	"time"
)

// wavenoteUser表 结构体  WavenoteUser
type WavenoteUser struct {
	Id           *int32     `json:"id" form:"id" gorm:"primarykey;column:id;"`                                           //id字段
	Email        *string    `json:"email" form:"email" gorm:"comment:邮箱;column:email;size:100;"`                         //邮箱
	ThirdAccount *string    `json:"thirdAccount" form:"thirdAccount" gorm:"comment:第三方账号;column:third_account;size:50;"` //第三方账号
	ChannelUid   *string    `json:"channelUid" form:"channelUid" gorm:"comment:渠道uid;column:channel_uid;size:100;"`      //渠道uid
	Pwd          *string    `json:"pwd" form:"pwd" gorm:"comment:密码;column:pwd;size:100;"`                               //密码
	Nickname     *string    `json:"nickname" form:"nickname" gorm:"comment:用户昵称;column:nickname;size:32;"`               //用户昵称
	Icon         *string    `json:"icon" form:"icon" gorm:"comment:头像;column:icon;size:100;"`                            //头像
	UserType     *int32     `json:"userType" form:"userType" gorm:"comment:用户类型;0普通用户;1VIP;column:user_type;"`           //用户类型;0普通用户;1VIP
	VipExpiredAt *int32     `json:"vipExpiredAt" form:"vipExpiredAt" gorm:"comment:vip用户的过期时间;column:vip_expired_at;"`   //vip用户的过期时间
	Status       *int32     `json:"status" form:"status" gorm:"comment:0-正常，1-锁定;column:status;"`                        //0-正常，1-锁定
	Origin       *int32     `json:"origin" form:"origin" gorm:"comment:注册来源:1邮箱注册;2 goole id;3 apple id;column:origin;"` //注册来源:1邮箱注册;2 goole id;3 apple id
	Lan          *string    `json:"lan" form:"lan" gorm:"comment:注册时候的系统语言;column:lan;size:20;"`                         //注册时候的系统语言
	TransLan     *string    `json:"transLan" form:"transLan" gorm:"comment:转写语言;column:trans_lan;size:20;"`              //转写语言
	Country      *string    `json:"country" form:"country" gorm:"comment:注册国家/地区编码;column:country;size:20;"`             //注册国家/地区编码
	Industry     *int32     `json:"industry" form:"industry" gorm:"comment:用户所属行业;column:industry;"`                     //用户所属行业
	DeviceId     *string    `json:"deviceId" form:"deviceId" gorm:"comment:设备ID;column:device_id;size:100;"`             //设备ID
	Ip           *string    `json:"ip" form:"ip" gorm:"comment:注册ip;column:ip;size:20;"`                                 //注册ip
	SyncStatus   *int32     `json:"syncStatus" form:"syncStatus" gorm:"comment:用户迁移状态:1迁移;2被迁移;column:sync_status;"`     //用户迁移状态:1迁移;2被迁移
	SyncUid      *int32     `json:"syncUid" form:"syncUid" gorm:"comment:迁移的用户uid;column:sync_uid;"`                     //迁移的用户uid
	CreatedAt    *time.Time `json:"createdAt" form:"createdAt" gorm:"comment:创建时间;column:created_at;"`                   //创建时间
	UpdatedAt    *time.Time `json:"updatedAt" form:"updatedAt" gorm:"comment:更新时间;column:updated_at;"`                   //更新时间
	DeletedAt    *int32     `json:"deletedAt" form:"deletedAt" gorm:"comment:删除写入时间戳;column:deleted_at;"`                //删除写入时间戳
}

// TableName wavenoteUser表 WavenoteUser自定义表名 user
func (WavenoteUser) TableName() string {
	return "user"
}
