package templates

const (
	GoModule = `module %s

go 1.17

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/stretchr/testify v1.4.0
)

`
)