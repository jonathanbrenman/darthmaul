package templates

const (
	AppFile = `package app

import (
	"%s/cmd/api/router"
)

func Start() {

	// Configure router
	router := router.NewRouter(":8080")
	router.Setup()
}

`
)
