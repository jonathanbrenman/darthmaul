package templates

const (
	MainTemplate = `package main

import "%s/cmd/api/app"

func main() {
	app.Start()
}
`
)
