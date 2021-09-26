package templates

const (
	DockerCompose = `version: '3.7'

services:
  %s-api:
    build:
      context: .
      dockerfile: dev.Dockerfile
    volumes:
      - ./api:/go/src/%s/api
    environment:
      - SCOPE=local
    ports:
      - "8080:8080"

`
)