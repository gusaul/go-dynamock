package dynamock

import (
	"github.com/pkg/errors"
)

var (
	ErrInvalidInput             = errors.New("invalid input")
	ErrNoExpectation            = errors.New("expectations not found")
	ErrNoTable                  = errors.New("expectations table not found")
	ErrNoKey                    = errors.New("expectations key not found")
	ErrNoItem                   = errors.New("expectations item not found")
	ErrTableExpectationMismatch = errors.New("expected table was not matched")
	ErrKeyExpectationMismatch   = errors.New("expected key was not matched")
	ErrItemExpectationMismatch  = errors.New("expected item was not matched")
)
