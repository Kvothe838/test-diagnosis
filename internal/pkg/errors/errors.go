package errors

import "github.com/pkg/errors"

var (
	PatientNotFoundErr = errors.New("patient not found")
)
