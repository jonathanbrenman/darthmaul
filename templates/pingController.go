package templates

const (
	PingControllerTemplate = `package controllers

import (
	"github.com/gin-gonic/gin"
)

type PingController interface {
	Ping(context *gin.Context)
}

type pingControllerImpl struct{}

func NewPingController() PingController {
	return &pingControllerImpl{}
}

func (pingCtrl *pingControllerImpl) Ping(c *gin.Context) {

	c.JSON(200, gin.H{
		"code":         200,
		"message":      "pong",
	})
}
`
	PingControllerTest = `package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	assert := assert.New(t)
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, r := gin.CreateTestContext(w)
	r.GET("/ping", NewPingController().Ping)
	c.Request, _ = http.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(w, c.Request)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}
	var response struct {
		Code        int    `+"`json:\"code\"`"+`
		Message     string `+"`json:\"message\"`"+`
	}

	json.NewDecoder(w.Body).Decode(&response)
	assert.Equal(200, response.Code)
	assert.Equal("pong", response.Message)
}
`
)
