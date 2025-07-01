package app_errors

import "errors"

var (
	ErrPackageNotFound      = errors.New("package not found")
	ErrNoCarrierFound       = errors.New("no carrier found")
	ErrInvalidPackageStatus = errors.New("package is not in a valid state to hire a carrier")
	ErrUserNotFound         = errors.New("user not found")
	ErrUserEmailExists      = errors.New("user with this email already exists")
)
