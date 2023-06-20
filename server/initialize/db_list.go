package initialize

import (
	"gorm.io/gorm"

	"github.com/pemako/gva/server/config"
	"github.com/pemako/gva/server/global"
)

const sys = "system"

func DBList() {
	dbMap := make(map[string]*gorm.DB)
	for _, info := range global.GvaConfig.DBList {
		if info.Disable {
			continue
		}
		switch info.Type {
		case "mysql":
			dbMap[info.AliasName] = GormMysqlByConfig(config.Mysql{GeneralDB: info.GeneralDB})
		case "mssql":
			dbMap[info.AliasName] = GormMssqlByConfig(config.Mssql{GeneralDB: info.GeneralDB})
		case "pgsql":
			dbMap[info.AliasName] = GormPgSqlByConfig(config.Pgsql{GeneralDB: info.GeneralDB})
		case "oracle":
			dbMap[info.AliasName] = GormOracleByConfig(config.Oracle{GeneralDB: info.GeneralDB})
		default:
			continue
		}
	}
	// 做特殊判断,是否有迁移
	// 适配低版本迁移多数据库版本
	if sysDB, ok := dbMap[sys]; ok {
		global.GvaDb = sysDB
	}
	global.GvaDbList = dbMap
}