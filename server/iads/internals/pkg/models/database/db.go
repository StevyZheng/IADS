package database

import (
	"github.com/jinzhu/gorm"
	mysql2 "iads/server/iads/internals/pkg/models/database/mysql"
	pgsql2 "iads/server/iads/internals/pkg/models/database/pgsql"
	config2 "iads/server/iads/pkg/config"
)

var DBE *gorm.DB

func CreateDBEngine() {
	dbType := config2.ConfValue.DBType
	if dbType == "pgsql" {
		pgsql2.InitPgsqlDB()
		DBE = pgsql2.Eloquent
	} else if dbType == "mysql" {
		mysql2.InitMysqlDB()
		DBE = mysql2.Eloquent
	}
}
