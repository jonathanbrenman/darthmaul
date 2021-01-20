package commands

type Command interface {
	Execute() (err error)
}