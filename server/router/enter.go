package router

import (
	"github.com/pemako/gva/server/router/example"
	"github.com/pemako/gva/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
