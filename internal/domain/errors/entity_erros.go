package errors

import (
	"errors"
	"strings"
	"sync"
)

type EntityErrors struct {
	mu     sync.Mutex
	errors []error
}

func (e *EntityErrors) Add(err string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.errors = append(e.errors, errors.New(err))
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

func (e *EntityErrors) ListErrors() string {
	e.mu.Lock()
	errs := e.errors
	e.mu.Unlock()
	var sb strings.Builder
	for _, err := range errs {
		sb.WriteString(err.Error())
		sb.WriteString("\n")
	}
	return sb.String()
}
