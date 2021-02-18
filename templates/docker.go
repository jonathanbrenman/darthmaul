package templates

const (
	DockerDev = `FROM golang:1.14-alpine
RUN apk add git gcc g++
ENV CGO_ENABLED 1
ENV GOPATH /go
ENV CC gcc


ADD . /go/src/%s
RUN cd /go/src/%s && go mod tidy
WORKDIR /go/src/%s/cmd/api

RUN go get -u github.com/cosmtrek/air
ENTRYPOINT  air run main.go
EXPOSE 8080

`
)