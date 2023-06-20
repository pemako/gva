package initialize

import (
	"github.com/songzhibin97/gkit/cache/local_cache"

	"github.com/pemako/gva/server/global"
	"github.com/pemako/gva/server/utils"
)

func OtherInit() {
	dr, err := utils.ParseDuration(global.GvaConfig.JWT.ExpiresTime)
	if err != nil {
		panic(err)
	}
	_, err = utils.ParseDuration(global.GvaConfig.JWT.BufferTime)
	if err != nil {
		panic(err)
	}

	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(dr),
	)
}
