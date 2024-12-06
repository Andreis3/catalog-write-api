package aggregates

import (
	"github.com/andreis3/catalog-write-api/internal/domain/commons"
	"github.com/andreis3/catalog-write-api/internal/domain/entities"
)

type CreateProduct struct {
	Product    *entities.Product
	Categories []*entities.Category
	Skus       []CreateSkus
}

type CreateSkus struct {
	Skus           *entities.Sku
	Medias         []*entities.Media
	Tags           []*entities.Tag
	Specifications []CreateSpecifications
	Offers         []CreateOffers
}

type CreateSpecifications struct {
	SpecificationKey   entities.SpecificationKey
	SpecificationValue entities.SpecificationValue
}

type CreateOffers struct {
	Offers       *entities.Offer
	Coordinate   *entities.Coordinate
	Installments []*entities.Installment
}

func CreateProductBuilder() *CreateProduct {
	return &CreateProduct{}
}

func (c *CreateProduct) WithProduct(product *entities.Product) *CreateProduct {
	c.Product = product
	return c
}

func (c *CreateProduct) WithCategories(categories []*entities.Category) *CreateProduct {
	c.Categories = categories
	return c
}

func (c *CreateProduct) WithSkus(skus []CreateSkus) *CreateProduct {
	c.Skus = skus
	return c
}

func (c *CreateProduct) Build() (*CreateProduct, *commons.EntityErrors) {
	product := &CreateProduct{
		Product:    c.Product,
		Categories: c.Categories,
		Skus:       c.Skus,
	}
	if err := product.Validate(); err != nil {
		return nil, err
	}
	return product, nil
}

func validate[T interface{ Validate() *commons.EntityErrors }](
	items []T,
	fieldName string,
	parentValidation *commons.EntityErrors,
) {
	var validationFields commons.ValidateFields
	if len(items) == 0 {
		parentValidation.Add(validationFields.CheckMinimumOfOne(len(items), fieldName))
		return
	}
	for i, item := range items {
		if validation := item.Validate(); validation != nil {
			parentValidation.MergeSlice(i, fieldName, validation)
		}
	}
}

func (c *CreateProduct) Validate() *commons.EntityErrors {
	validation := c.Product.EntityErrors
	validationFields := c.Product.ValidateFields

	if productValidation := c.Product.Validate(); productValidation != nil {
		validation.Merge("product", productValidation)
	}
	validate(c.Categories, "categories", &validation)
	if c.checkCircularDependenciesCategories(c.Categories) {
		validation.Add(validationFields.CheckCircularDependencies("categories"))
	}
	if len(c.Skus) == 0 {
		validation.Add(validationFields.CheckMinimumOfOne(len(c.Skus), "skus"))
		return &validation
	}

	for i, sku := range c.Skus {
		skuValidation := sku.Skus.Validate()

		validate(sku.Medias, "medias", skuValidation)
		validate(sku.Tags, "tags", skuValidation)
		validateSpecifications(sku.Specifications, skuValidation)
		validateOffers(sku.Offers, skuValidation)
		validation.MergeSlice(i, "skus", skuValidation)
	}
	return &validation
}

func validateSpecifications(specs []CreateSpecifications, validation *commons.EntityErrors) {
	for i, spec := range specs {
		var specValidation *commons.EntityErrors
		if keyValidation := spec.SpecificationKey.Validate(); keyValidation != nil {
			specValidation.Merge("key", keyValidation)
		}
		if valueValidation := spec.SpecificationValue.Validate(); valueValidation != nil {
			specValidation.Merge("value", valueValidation)
		}
		validation.MergeSlice(i, "specifications", specValidation)
	}
}

func validateOffers(offers []CreateOffers, validation *commons.EntityErrors) {
	for i, offer := range offers {
		offerValidation := offer.Offers.Validate()

		validate(offer.Installments, "installments", offerValidation)
		if coordValidation := offer.Coordinate.Validate(); coordValidation != nil {
			offerValidation.Merge("coordinate", coordValidation)
		}
		validation.MergeSlice(i, "offers", offerValidation)
	}
}

func (c *CreateProduct) checkCircularDependenciesCategories(categories []*entities.Category) bool {
	grafo := make(map[string][]string)
	inDegree := make(map[string]int)

	for _, category := range categories {
		if category.GetParentCategoryKey() != "" {
			grafo[category.GetParentCategoryKey()] = append(grafo[category.GetParentCategoryKey()], category.GetCategoryKey())
			inDegree[category.GetCategoryKey()]++
		}
		if _, exists := grafo[category.GetCategoryKey()]; !exists {
			inDegree[category.GetCategoryKey()] = 0
		}
	}

	var queue []string
	for no, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, no)
		}
	}

	visited := 0
	for len(queue) > 0 {
		no := queue[0]
		queue = queue[1:]
		visited++
		for _, neighbor := range grafo[no] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}
	return visited != len(inDegree)
}
