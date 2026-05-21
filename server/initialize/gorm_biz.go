package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate()
	if err != nil {
		return err
	}
	ai_userDb := global.GetGlobalDBByDBName("ai_user")
	ai_userDb.AutoMigrate(system.AudioPpt{})
	return nil
}
