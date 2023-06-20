package response

import "github.com/pemako/gva/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
