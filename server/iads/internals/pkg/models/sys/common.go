package sys

import (
	database2 "iads/server/iads/internals/pkg/models/database"
	config2 "iads/server/iads/pkg/config"
)

func ModelInit() {
	database2.CreateDBEngine()
	if config2.ConfValue.DBType == "mysql" {
		database2.DBE.Set("gorm:table_options", "DEFAULT CHARSET=utf8 AUTO_INCREMENT=1").AutoMigrate(&Role{}, &User{})
	} else if config2.ConfValue.DBType == "pgsql" {
		database2.DBE.AutoMigrate(&Role{}, &User{})
	}
	database2.DBE.Model(&User{}).AddForeignKey("role_id", "roles(id)", "RESTRICT", "RESTRICT")
}
