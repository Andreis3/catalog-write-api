package errors

import (
	"strings"
	"sync"
)

type EntityErrors struct {
	mu     sync.Mutex
	errors []error
}

func (e *EntityErrors) Add(err error) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.errors = append(e.errors, err)
}

func (e *EntityErrors) HasErrors() bool {
	e.mu.Lock()
	defer e.mu.Unlock()
	return len(e.errors) > 0
}

func (e *EntityErrors) Errors() []error {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.errors
}

func (e *EntityErrors) Error() string {
	e.mu.Lock()
	defer e.mu.Unlock()
	return strings.Join(e.ErrorsToString(), "\n")
}

func (e *EntityErrors) ErrorsToString() []string {
	errString := make([]string, len(e.errors))
	for i, err := range e.errors {
		errString[i] = err.Error()
	}
	return errString
}
