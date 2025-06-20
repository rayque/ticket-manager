package app_errors

import "errors"

var (
	ErrPackageNotFound = errors.New("package not found")
	ErrNoCarrierFound  = errors.New("no carrier found")
)
