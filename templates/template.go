package templates

const (
	Boilerplate = `package template

type Darth interface {
	MyMethod() (err error)
}

type darthImpl struct {}

var singletonDarth *darthImpl

func NewDarth() Darth {
	if singletonDarth != nil {
		return singletonDarth
	}

	singletonDarth = &darthImpl{}
	return singletonDarth
}

func (%s *darthImpl) MyMethod() (err error) {
	// Here put your code
	return err
}
`
)
