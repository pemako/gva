package service

import (
	"github.com/pemako/gva/server/service/example"
	"github.com/pemako/gva/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
