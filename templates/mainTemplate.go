package templates

const (
	MainTemplate = `package main

func main() {
	// Configure router
	router := router.NewRouter(":8080")
	router.Setup()
}
`
)
