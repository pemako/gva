package system

import (
	"github.com/pemako/gva/server/config"
)

// System 配置文件结构体
type System struct {
	Config config.Server `json:"config"`
}
