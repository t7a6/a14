package a14

import (
	"fmt"
	"golang.org/x/xerrors"
)

type ErrorKind string

const (
	// UnauthorizedError This error is used when the user is not authorized to access the resource.
	UnauthorizedError ErrorKind = "unauthorized_error"
	// ValidationError This error is used when the request data is invalid.
	ValidationError ErrorKind = "validation_error"
	// DuplicateError This error is used when the data already exists.
	DuplicateError ErrorKind = "duplicate_error"
	// NotFoundError This error is used when the data does not exist.
	NotFoundError ErrorKind = "not_found_error"
	// NetworkError This error is used when a network error occurs.
	NetworkError ErrorKind = "network_error"
	// ExternalApiError This error is used when an error occurs in the external API.
	ExternalApiError ErrorKind = "external_api_error"
	// InternalApiError This error is used when an internal error occurs in the internal API.
	InternalApiError ErrorKind = "internal_api_error"
)

var _ Error = (*errorInternal)(nil)

type Error interface {
	Error() string
	Kind() ErrorKind
}

type errorInternal struct {
	kind  ErrorKind
	msg   string
	err   error
	frame xerrors.Frame
}

func (e *errorInternal) Error() string {
	return e.msg
}

func (e *errorInternal) Kind() ErrorKind {
	return e.kind
}

func (e *errorInternal) Format(s fmt.State, v rune) {
	xerrors.FormatError(e, s, v)
}

func (e *errorInternal) FormatError(p xerrors.Printer) (next error) {
	p.Print(fmt.Sprintf("[%s] %s", e.kind, e.msg))
	e.frame.Format(p)
	return e.err
}

func NewError(msg string, kind ErrorKind, err error) Error {
	return &errorInternal{
		kind:  kind,
		msg:   msg,
		err:   err,
		frame: xerrors.Caller(1),
	}
}
