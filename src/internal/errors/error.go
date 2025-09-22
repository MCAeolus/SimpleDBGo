package errors

import "errors"

var (
	// field
	ErrIllegalCast = errors.New("field was not able to be cast as expected")
	ErrOpComparisonNotImplemented = errors.New("comparison with op not implemented for field")

	// tuple
	ErrDoesNotExist = errors.New("requested field or index does not exist")
)