package global

import (
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"

	"github.com/pemako/gva/server/config"
	"github.com/pemako/gva/server/utils/timer"
)

var (
	GvaDb                 *gorm.DB
	GvaDbList             map[string]*gorm.DB
	GvaRedis              *redis.Client
	GvaConfig             config.Server
	GvaViper              *viper.Viper
	GvaLog                *zap.Logger
	GvaTimer              = timer.NewTimerTask()
	GvaConcurrencyControl = &singleflight.Group{}

	BlackCache local_cache.Cache
	lock       sync.RWMutex
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return GvaDbList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := GvaDbList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}
