package entities

import (
	"github.com/andreis3/catalog-write-api/internal/domain/commons"
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
	commons.EntityErrors
	commons.ValidateFields
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

func (a *APIKey) Validate() *commons.EntityErrors {
	a.Add(a.CheckEmptyField(a.name, "name"))
	a.Add(a.CheckMinCharacters(a.name, "name", 3))
	a.Add(a.CheckEmptyField(a.status, "status"))
	a.Add(a.CheckIsValidStatus(a.status, "status", APIKeyStatus[:]))
	return &a.EntityErrors
}
