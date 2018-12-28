package api

import "errors"

var (
	// ErrUnknowed ..
	ErrUnknowed = errors.New("UnknowedError")
	// ErrPanicNotNilNotError ...
	ErrPanicNotNilNotError = errors.New("Panic not nil not error")
)

// Error ...
type Error struct {
	Status   int
	Location string
	error
}
