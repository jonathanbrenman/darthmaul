package templates

const (
	DockerCompose = `version: '3.7'

services:
  %s-api:
    build:
      context: .
      dockerfile: dev.Dockerfile
    volumes:
      - ./cmd/api:/go/src/%s/cmd/api
    environment:
      - SCOPE=local
    ports:
      - "8080:8080"

`
)