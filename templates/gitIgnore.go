package templates

const (
	GitIgnore = `### Go ###
# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

*.test

# Output of the go profiling tool
*.prof

# Output of the go coverage tool
*.out

# Compiled Object files, Static and Dynamic libs (Shared Objects)
*.o
*.a

# Go generated folders
_obj
_test
vendor

# Architecture specific extensions/prefixes
*.[568vq]

# CGO specific files
*.cgo1.go
*.cgo2.c
_cgo_defun.c
_cgo_gotypes.go
_cgo_export.*

# Common IDE folders
.idea
.vscode
**/tmp
api/tmp/*
**/tmp
`
)
