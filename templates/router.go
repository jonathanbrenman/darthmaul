package templates

const (
	UrlsMapping = `package router

import "%s/cmd/api/router/factory"

func (r routerImpl) routes() {
	factoryCtrl := factory.NewCtrlFactory()

	r.router.GET("/ping", factoryCtrl.BuildPingController().Ping)
}
`
	RouterTest = `package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCORSMiddleware(t *testing.T) {
	router := NewRouter(":8080")
	go router.Setup()

	origin := "*"
	server := httptest.NewServer(router.GetRouter())
	defer server.Close()

	client := &http.Client{}

	// TESTING CORS WITH OPTIONS METHOD
	req, _ := http.NewRequest(
		"OPTIONS",
		"http://localhost:8080/ping",
		nil,
	)
	req.Header.Add("Origin", origin)

	get, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	o := get.Header.Get("Access-Control-Allow-Origin")
	if o != origin {
		t.Errorf("Got '%s' ; expecting origin '%s'", o, origin)
	}

	// TESTING CORS WITH GET METHOD
	req, _ = http.NewRequest(
		"GET",
		"http://localhost:8080/ping",
		nil,
	)
	req.Header.Add("Origin", origin)

	get, err = client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	o = get.Header.Get("Access-Control-Allow-Origin")
	if o != origin {
		t.Errorf("Got '%s' ; expecting origin '%s'", o, origin)
	}
}
`
	RouterTemplate = `package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"%s/cmd/api/middlewares"
)

type routerImpl struct {
	router *gin.Engine
	port   string
}

type Router interface {
	Setup()
	GetRouter() *gin.Engine
}

func NewRouter(port string) *routerImpl {
	return &routerImpl{
		router: gin.Default(),
		port:   port,
	}
}

// Router Setup
func (r routerImpl) configure() {
	r.router.Use(middlewares.CORSMiddleware())
	r.routes()
	if err := r.router.Run(r.port); err != nil {
		fmt.Errorf("Unable to start router error: %v", err)
		panic(err)
	}
}

func (r routerImpl) Setup() {
	r.configure()
}

func (r routerImpl) GetRouter() *gin.Engine {
	return r.router
}
`
)
