package entities

import "github.com/andreis3/catalog-write-api/internal/domain/commons"

type Tag struct {
	id    int64
	skuID int64
	name  string
	commons.EntityErrors
	commons.ValidateFields
}

func TagBuilder() *Tag {
	return &Tag{}
}

func (t *Tag) GetID() int64 {
	return t.id
}

func (t *Tag) GetSkuID() int64 {
	return t.skuID
}

func (t *Tag) GetName() string {
	return t.name
}

func (t *Tag) SetID(id int64) *Tag {
	t.id = id
	return t
}

func (t *Tag) SetSkuID(skuID int64) *Tag {
	t.skuID = skuID
	return t
}

func (t *Tag) SetName(name string) *Tag {
	t.name = name
	return t
}

func (t *Tag) Build() *Tag {
	return t
}

func (t *Tag) Validate() *commons.EntityErrors {
	t.Add(t.CheckEmptyField(t.name, "name"))
	t.Add(t.CheckMaxCharacters(t.name, "name", 10))
	return &t.EntityErrors
}
