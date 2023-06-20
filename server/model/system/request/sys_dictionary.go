package request

import (
	"github.com/pemako/gva/server/model/common/request"
	"github.com/pemako/gva/server/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
