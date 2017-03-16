package arango

import "github.com/pkg/errors"

type ErrorCode string

func (e ErrorCode) String() string {
	return string(e)
}

func (e ErrorCode) Error() error {
	return errors.New(e.String())
}

const (
	ErrorCodeNoDatabaseSelected ErrorCode = "no database selected"
)
