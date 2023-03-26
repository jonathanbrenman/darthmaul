package templates

const (
	ControllerFactory = `package controllers

type ControllerBuilder interface {
	BuildPingController() PingController
}

type controllerBuildImpl struct {}

func NewCtrlFactory() ControllerBuilder {
	return &controllerBuildImpl{}
}

func (ctrlFactory *controllerBuildImpl) BuildPingController() PingController {
	return NewPingController()
}
`
	FactoryTest = `package controllers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildPingController(t *testing.T) {
	factoryCtrl := NewCtrlFactory()
	pingCtrl := factoryCtrl.BuildPingController()
	assert.NotNil(t, pingCtrl)
}`
)
