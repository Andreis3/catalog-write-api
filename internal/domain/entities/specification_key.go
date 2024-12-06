package entities

import "github.com/andreis3/catalog-write-api/internal/domain/commons"

type SpecificationKey struct {
	id       int64
	apiKeyID int64
	key      string
	commons.EntityErrors
	commons.ValidateFields
}

func SpecificationKeyBuilder() *SpecificationKey {
	return &SpecificationKey{}
}

func (s *SpecificationKey) GetID() int64 {
	return s.id
}

func (s *SpecificationKey) GetKey() string {
	return s.key
}

func (s *SpecificationKey) GetAPIKeyID() int64 {
	return s.apiKeyID
}

func (s *SpecificationKey) SetID(id int64) *SpecificationKey {
	s.id = id
	return s
}

func (s *SpecificationKey) SetKey(key string) *SpecificationKey {
	s.key = key
	return s
}

func (s *SpecificationKey) SetAPIKeyID(apiKeyID int64) *SpecificationKey {
	s.apiKeyID = apiKeyID
	return s
}

func (s *SpecificationKey) Build() *SpecificationKey {
	return s
}

func (s *SpecificationKey) Validate() *commons.EntityErrors {
	s.Add(s.CheckEmptyField(s.key, "key"))
	s.Add(s.CheckMaxCharacters(s.key, "key", 10))
	return &s.EntityErrors
}
