package initialiaze

import (
	"admin/global"
	"admin/models"
)

func Migrate() {
	global.DB.AutoMigrate(
		models.AdminUser{},
		models.AdminRole{},
		models.AdminMenu{},
	)
}
