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
	if err == nil {
		return
	}
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

func (e *EntityErrors) ListErrors() string {
	e.mu.Lock()
	defer e.mu.Unlock()
	var errs []string
	for _, err := range e.errors {
		errs = append(errs, err.Error())
	}
	return strings.Join(errs, "\n")
}
