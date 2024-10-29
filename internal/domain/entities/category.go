package entities

import "github.com/andreis3/catalog-write-api/internal/domain/errors"

type Category struct {
	id                int64
	apiKeyID          int64
	parentID          int64
	categoryKey       string
	description       string
	parentCategoryKey string
	errors.EntityErrors
}

func CategoryBuilder() *Category {
	return &Category{}
}

func (c *Category) GetID() int64 {
	return c.id
}

func (c *Category) GetCategoryKey() string {
	return c.categoryKey
}

func (c *Category) GetDescription() string {
	return c.description
}

func (c *Category) GetParentID() int64 {
	return c.parentID
}

func (c *Category) GetParentCategoryKey() string {
	return c.parentCategoryKey
}

func (c *Category) GetAPIKeyID() int64 {
	return c.apiKeyID
}

func (c *Category) SetID(id int64) *Category {
	c.id = id
	return c
}

func (c *Category) SetCategoryKey(categoryKey string) *Category {
	c.categoryKey = categoryKey
	return c
}

func (c *Category) SetDescription(description string) *Category {
	c.description = description
	return c
}

func (c *Category) SetParentID(parentID int64) *Category {
	c.parentID = parentID
	return c
}

func (c *Category) SetParentCategoryKey(parentCategoryKey string) *Category {
	c.parentCategoryKey = parentCategoryKey
	return c
}

func (c *Category) SetAPIKeyID(apiKeyID int64) *Category {
	c.apiKeyID = apiKeyID
	return c
}

func (c *Category) Build() *Category {
	return c
}

func (c *Category) Validate() *errors.EntityErrors {
	if c.categoryKey == "" {
		c.Add("category_key: is required")
	}

	if c.description == "" {
		c.Add("description: is required")
	}

	return &c.EntityErrors
}
