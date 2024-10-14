package entities

import (
	"errors"
	"slices"

	err "github.com/andreis3/catalog-write-api/internal/domain/errors"
)

const (
	ACTIVE   = "active"
	INACTIVE = "inactive"
)

var APIKeyStatus = [...]string{ACTIVE, INACTIVE}

type APIKey struct {
	ID     int64
	Name   string
	Status string
	err.EntityErrors
}

func Builder() *APIKey {
	return &APIKey{}
}

func (a *APIKey) SetID(id int64) *APIKey {
	a.ID = id
	return a
}

func (a *APIKey) SetName(name string) *APIKey {
	a.Name = name
	return a
}

func (a *APIKey) SetStatus(status string) *APIKey {
	a.Status = status
	return a
}

func (a *APIKey) SetStatusActive() *APIKey {
	a.Status = ACTIVE
	return a
}

func (a *APIKey) SetStatusInactive() *APIKey {
	a.Status = INACTIVE
	return a
}

func (a *APIKey) Build() *APIKey {
	return a
}

func (a *APIKey) Validate() *err.EntityErrors {
	if a.Name == "" {
		a.Add(errors.New("name: is required"))
	}
	if a.Status == "" {
		a.Add(errors.New("status: is required"))
	}
	if a.Status != "" && !slices.Contains(APIKeyStatus[:], a.Status) {
		a.Add(errors.New("status: is invalid, valid values are active or inactive"))
	}
	return &a.EntityErrors
}
