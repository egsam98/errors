package example2

import (
	stderrors "errors"

	"github.com/egsam98/errors"
)

var ErrTest = stderrors.New("test")

func Method() error {
	return errors.Wrap(ErrTest, "test")
}
