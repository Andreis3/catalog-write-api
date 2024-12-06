package entities

import "github.com/andreis3/catalog-write-api/internal/domain/commons"

type SpecificationValue struct {
	id                 int64
	specificationKeyID int64
	value              string
	commons.EntityErrors
	commons.ValidateFields
}

func SpecificationValueBuilder() *SpecificationValue {
	return &SpecificationValue{}
}

func (s *SpecificationValue) GetID() int64 {
	return s.id
}

func (s *SpecificationValue) GetSpecificationKeyID() int64 {
	return s.specificationKeyID
}

func (s *SpecificationValue) GetValue() string {
	return s.value
}

func (s *SpecificationValue) SetID(id int64) *SpecificationValue {
	s.id = id
	return s
}

func (s *SpecificationValue) SetSpecificationKeyID(specificationKeyID int64) *SpecificationValue {
	s.specificationKeyID = specificationKeyID
	return s
}

func (s *SpecificationValue) SetValue(value string) *SpecificationValue {
	s.value = value
	return s
}

func (s *SpecificationValue) Build() *SpecificationValue {
	return s
}

func (s *SpecificationValue) Validate() *commons.EntityErrors {
	s.Add(s.CheckEmptyField(s.value, "value"))
	s.Add(s.CheckMaxCharacters(s.value, "value", 10))
	return &s.EntityErrors
}
