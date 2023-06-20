package main

import (
	"go.uber.org/zap"

	"github.com/pemako/gva/server/core"
	"github.com/pemako/gva/server/global"
	"github.com/pemako/gva/server/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title                       Swagger Example API
// @version                     0.0.1
// @description                 This is a sample Server pets
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	global.GvaViper = core.Viper() // 初始化Viper
	initialize.OtherInit()
	global.GvaLog = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GvaLog)
	global.GvaDb = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	initialize.DBList()
	if global.GvaDb != nil {
		initialize.RegisterTables() // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GvaDb.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
