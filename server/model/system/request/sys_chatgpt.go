package request

import (
	"github.com/pemako/gva/server/model/common/request"
	"github.com/pemako/gva/server/model/system"
)

type ChatGptRequest struct {
	system.ChatGpt
	request.PageInfo
}
