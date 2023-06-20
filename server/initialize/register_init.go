package initialize

import (
	_ "github.com/pemako/gva/server/source/example"
	_ "github.com/pemako/gva/server/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
