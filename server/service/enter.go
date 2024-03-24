package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/step"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/zeppLife"
)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	ExampleServiceGroup  example.ServiceGroup
	StepServiceGroup     step.ServiceGroup
	ZeppLifeServiceGroup zeppLife.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
