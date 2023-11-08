package finaplan

import (
	"errors"
	"fmt"
)

var ErrIntervalsLessThanOne = errors.New("intervals must be equal or greater than 1")

type ErrDecodeDecimal struct {
	paramName string
	err       error
}

func NewErrDecodeDecimal(paramName string, err error) ErrDecodeDecimal {
	return ErrDecodeDecimal{
		paramName: paramName,
		err:       err,
	}
}

func (e ErrDecodeDecimal) Error() string {
	return fmt.Sprintf("decoding %s as decimal: %s", e.paramName, e.err.Error())
}

func (e ErrDecodeDecimal) Unwrap() error {
	return e.err
}
