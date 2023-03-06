package example

import (
	"github.com/egsam98/errors"
	"github.com/egsam98/errors/example/example2"
)

type Struct struct{}

func (s Struct) Method() error {
	err := example2.Method()
	return errors.Wrap(err, "prefix")
}
