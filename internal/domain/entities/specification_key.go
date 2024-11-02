package entities

import "github.com/andreis3/catalog-write-api/internal/domain/errors"

type SpecificationKey struct {
	id  int64
	key string
	errors.EntityErrors
	errors.ValidateFields
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

func (s *SpecificationKey) SetID(id int64) *SpecificationKey {
	s.id = id
	return s
}

func (s *SpecificationKey) SetKey(key string) *SpecificationKey {
	s.key = key
	return s
}

func (s *SpecificationKey) Build() *SpecificationKey {
	return s
}

func (s *SpecificationKey) Validate() *errors.EntityErrors {
	s.Add(s.CheckEmptyField(s.key, "key"))
	s.Add(s.CheckMaxCharacters(s.key, "key", 10))
	return &s.EntityErrors
}
