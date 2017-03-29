package errc

import "github.com/pkg/errors"

type ErrorCoder interface {
	ErrorCode() ErrorCode
}

type ErrorCode string

func (e ErrorCode) String() string {
	return string(e)
}

func (e ErrorCode) Error() error {
	return errors.New(e.String())
}

func (e ErrorCode) Msg() string {
	return errorCodeMsg[e]
}

const (

	//String Checks
	ErrorCodeInvalidEdgeAttribute ErrorCode = "invalid edge attribute"

	//Database Errors
	ErrorCodeNoDatabaseSelected   ErrorCode = "no database selected"
	ErrorCodeCollectionNotExist   ErrorCode = "collection is left blank"
	ErrorInternalServer           ErrorCode = "Fatal Error"
)
