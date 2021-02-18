package templates

const (
	Readme = `# Service Template created by darthmaul cli
Basic scaffolding created with darthmaul cli

# Defaults
  - Port: [8080]
  - Gin Gonic
  - docker y docker-compose.

# Run the app
`+"```sh" + `
$ cd <app-name>
$ docker-compose up
` + "```"

)