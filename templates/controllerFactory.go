package templates

const (
	ControllerFactory = `package factory

import "%s/cmd/api/controllers"

type ControllerBuilder interface {
	BuildPingController() controllers.PingController
}

type controllerBuildImpl struct {}

func NewCtrlFactory() ControllerBuilder {
	return &controllerBuildImpl{}
}

func (ctrlFactory *controllerBuildImpl) BuildPingController() controllers.PingController {
	return controllers.NewPingController()
}
`
	FactoryTest = `package factory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildPingController(t *testing.T) {
	assert := assert.New(t)
	factoryCtrl := NewCtrlFactory()
	pingCtrl := factoryCtrl.BuildPingController()
	assert.NotNil(pingCtrl)
}`
)
