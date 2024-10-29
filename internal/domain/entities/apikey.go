package entities

import (
	"slices"

	"github.com/andreis3/catalog-write-api/internal/domain/errors"
)

const (
	ACTIVE   = "active"
	INACTIVE = "inactive"
)

var APIKeyStatus = [...]string{ACTIVE, INACTIVE}

type APIKey struct {
	id     int64
	name   string
	status string
	errors.EntityErrors
}

func ApiKeyBuilder() *APIKey {
	return &APIKey{}
}

func (a *APIKey) GetID() int64 {
	return a.id
}

func (a *APIKey) GetName() string {
	return a.name
}

func (a *APIKey) GetStatus() string {
	return a.status
}

func (a *APIKey) SetID(id int64) *APIKey {
	a.id = id
	return a
}

func (a *APIKey) SetName(name string) *APIKey {
	a.name = name
	return a
}

func (a *APIKey) SetStatus(status string) *APIKey {
	a.status = status
	return a
}

func (a *APIKey) Activate() *APIKey {
	a.status = ACTIVE
	return a
}

func (a *APIKey) Deactivate() *APIKey {
	a.status = INACTIVE
	return a
}

func (a *APIKey) Build() *APIKey {
	return a
}

func (a *APIKey) Validate() *errors.EntityErrors {
	if a.name == "" {
		a.Add("name: is required")
	} else if len(a.name) < 3 {
		a.Add("name: is too short, minimum length is 3 characters")
	}
	if a.status == "" {
		a.Add("status: is required")
	} else if !slices.Contains(APIKeyStatus[:], a.status) {
		a.Add("status: is invalid, valid values are active or inactive")
	}
	return &a.EntityErrors
}
