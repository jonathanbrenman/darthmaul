package templates

const (
	DockerDev = `FROM golang:1.18-alpine
RUN apk add git gcc g++
ENV CGO_ENABLED 1
ENV GOPATH /go
ENV CC gcc


ADD . /go/src/%s
RUN cd /go/src/%s && go mod tidy
WORKDIR /go/src/%s/api

RUN go install github.com/cosmtrek/air@latest
ENTRYPOINT air run main.go
EXPOSE 8080

`
)
