package templates

const (
	Boilerplate = `package template

type Darth interface {
	MyMethod() (err error)
}

type darthImpl struct {}

func NewDarth() Darth {
	return &darthImpl{}
}

func (d darthImpl) MyMethod() (err error) {
	// Here put your code
	return err
}
`
)