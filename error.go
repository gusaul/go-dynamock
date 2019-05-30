package dynamock

import (
	"github.com/pkg/errors"
)

var (
	ErrNoExpectation = errors.New("expectation not found")
	ErrTableExpectationMismatch = errors.New("expected table was not matched")
	ErrKeyExpectationMismatch = errors.New("expected key was not matched")
)
