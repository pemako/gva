package initialize

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/pemako/gva/server/global"
	"github.com/pemako/gva/server/middleware"
	"github.com/pemako/gva/server/plugin/email"
	"github.com/pemako/gva/server/utils/plugin"
)

func PluginInit(group *gin.RouterGroup, Plugin ...plugin.Plugin) {
	for i := range Plugin {
		PluginGroup := group.Group(Plugin[i].RouterPath())
		Plugin[i].Register(PluginGroup)
	}
}

func InstallPlugin(Router *gin.Engine) {
	PublicGroup := Router.Group("")
	fmt.Println("无鉴权插件安装==》", PublicGroup)
	PrivateGroup := Router.Group("")
	fmt.Println("鉴权插件安装==》", PrivateGroup)
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	//  添加跟角色挂钩权限的插件 示例 本地示例模式于在线仓库模式注意上方的import 可以自行切换 效果相同
	PluginInit(PrivateGroup, email.CreateEmailPlug(
		global.GvaConfig.Email.To,
		global.GvaConfig.Email.From,
		global.GvaConfig.Email.Host,
		global.GvaConfig.Email.Secret,
		global.GvaConfig.Email.Nickname,
		global.GvaConfig.Email.Port,
		global.GvaConfig.Email.IsSSL,
	))
}
