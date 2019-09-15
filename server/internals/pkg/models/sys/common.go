package sys

import (
	"iads/server/internals/pkg/models/database"
	"iads/server/pkg/config"
)

func ModelInit() {
	database.CreateDBEngine()
	if config.ConfValue.DBType == "mysql" {
		database.DBE.Set("gorm:table_options", "DEFAULT CHARSET=utf8 AUTO_INCREMENT=1").AutoMigrate(&Role{}, &User{})
	} else if config.ConfValue.DBType == "pgsql" {
		database.DBE.AutoMigrate(&Role{}, &User{})
	}
	database.DBE.Model(&User{}).AddForeignKey("role_id", "roles(id)", "RESTRICT", "RESTRICT")
}
