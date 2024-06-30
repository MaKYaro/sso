package storage

import "errors"

var (
	ErrUserExists   = errors.New("user already exists")
	ErrUserNotFound = errors.New("user isn't found")
	ErrAppNotFound  = errors.New("app isn't found")
)
