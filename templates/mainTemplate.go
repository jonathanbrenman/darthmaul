package templates

const (
	MainTemplate = `package main

import (
	"%s/api/router"
)

func main() {
	// Configure router
	router := router.NewRouter(":8080")
	router.Setup()
}
`
)
