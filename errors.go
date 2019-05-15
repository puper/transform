package transform

import (
	"errors"
)

var (
	ErrTypeNotMatch  = errors.New("type not match")
	ErrConvertFailed = errors.New("convert failed")
	ErrFieldNotFound = errors.New("field not found")
)

type ProcessError struct {
	Field string
	err   error
}

func NewProcessError(field string, err error) *ProcessError {
	return &ProcessError{
		Field: field,
		err:   err,
	}
}

func (this ProcessError) Error() string {
	return this.Field + ":" + this.err.Error()
}
