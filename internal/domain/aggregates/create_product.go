package aggregates

import (
	"github.com/andreis3/catalog-write-api/internal/domain/entities"
	"github.com/andreis3/catalog-write-api/internal/domain/errors"
)

type CreateProduct struct {
	Product    entities.Product
	Categories []entities.Category
	Skus       []CreateSkus
}

type CreateSkus struct {
	Skus           entities.Sku
	Medias         []entities.Media
	Tags           []entities.Tag
	Specifications []CreateSpecifications
	Offers         []CreateOffers
}

type CreateSpecifications struct {
	SpecificationKey   entities.SpecificationKey
	SpecificationValue entities.SpecificationValue
}

type CreateOffers struct {
	Offers       entities.Offer
	Coordinate   entities.Coordinate
	Installments []entities.Installment
}

func (c *CreateProduct) Validate() *errors.EntityErrors {
	productValidation := c.Product.Validate()
	productValidation.Add(c.Product.CheckMinimumOfOne(len(c.Categories), "categories"))
	c.validateCategories(productValidation)
	productValidation.Add(c.Product.CheckMinimumOfOne(len(c.Skus), "skus"))
	c.validateSkus(productValidation)
	return productValidation
}

func (c *CreateProduct) validateCategories(productValidation *errors.EntityErrors) {
	for index := range c.Categories {
		categoryValidation := c.Categories[index].Validate()
		productValidation.MergeSlice(index, "categories", categoryValidation)
	}
}

func (c *CreateProduct) validateSkus(productValidation *errors.EntityErrors) {
	var skuValidation *errors.EntityErrors
	for index := range c.Skus {
		skuValidation = c.Skus[index].Skus.Validate()
		c.validateMedias(c.Skus[index].Medias, skuValidation)
		c.validateTags(c.Skus[index].Tags, skuValidation)
		c.validateSpecifications(c.Skus[index].Specifications, skuValidation)
		c.validateOffers(c.Skus[index].Offers, skuValidation)
		productValidation.MergeSlice(index, "skus", skuValidation)
	}
}

func (c *CreateProduct) validateMedias(tags []entities.Media, skuValidation *errors.EntityErrors) {
	skuValidation.Add(c.Product.CheckMinimumOfOne(len(tags), "medias"))
	for index := range tags {
		mediaValidation := tags[index].Validate()
		skuValidation.MergeSlice(index, "medias", mediaValidation)
	}
}

func (c *CreateProduct) validateTags(tags []entities.Tag, skuValidation *errors.EntityErrors) {
	skuValidation.Add(c.Product.CheckMinimumOfOne(len(tags), "tags"))
	for index := range tags {
		tagValidation := tags[index].Validate()
		skuValidation.MergeSlice(index, "tags", tagValidation)
	}
}

func (c *CreateProduct) validateSpecifications(specifications []CreateSpecifications, skuValidation *errors.EntityErrors) {
	for index := range specifications {
		specificationKeyValidation := specifications[index].SpecificationKey.Validate()
		skuValidation.MergeSlice(index, "key", specificationKeyValidation)
		specificationValueValidation := specifications[index].SpecificationValue.Validate()
		skuValidation.MergeSlice(index, "value", specificationValueValidation)
	}
}

func (c *CreateProduct) validateOffers(offers []CreateOffers, skuValidation *errors.EntityErrors) {
	for index := range offers {
		offerValidation := offers[index].Offers.Validate()
		c.validateInstallments(offers[index].Installments, offerValidation)
		coordinateValidation := offers[index].Coordinate.Validate()
		offerValidation.Merge("coordinate", coordinateValidation)
		skuValidation.MergeSlice(index, "offers", offerValidation)
	}
}

func (c *CreateProduct) validateInstallments(installments []entities.Installment, offerValidation *errors.EntityErrors) {
	offerValidation.Add(c.Product.CheckMinimumOfOne(len(installments), "installments"))
	for index := range installments {
		installmentValidation := installments[index].Validate()
		offerValidation.MergeSlice(index, "installments", installmentValidation)
	}
}

func CreateProductBuilder() *CreateProduct {
	return &CreateProduct{}
}

func (c *CreateProduct) SetProduct(product entities.Product) *CreateProduct {
	c.Product = product
	return c
}

func (c *CreateProduct) SetCategories(categories []entities.Category) *CreateProduct {
	c.Categories = categories
	return c
}

func (c *CreateProduct) SetSkus(skus []CreateSkus) *CreateProduct {
	c.Skus = skus
	return c
}

func (c *CreateProduct) SetSkuTags(sku CreateSkus, tags []entities.Tag) *CreateProduct {
	sku.Tags = tags
	return c
}

func (c *CreateProduct) SetSkuMedias(sku CreateSkus, medias []entities.Media) *CreateProduct {
	sku.Medias = medias
	return c
}

func (c *CreateProduct) SetSkuSpecifications(sku CreateSkus, specifications []CreateSpecifications) *CreateProduct {
	sku.Specifications = specifications
	return c
}

func (c *CreateProduct) SetSkuOffers(sku CreateSkus, offers []CreateOffers) *CreateProduct {
	sku.Offers = offers
	return c
}

func (c *CreateProduct) SetOfferInstallments(offer CreateOffers, installments []entities.Installment) *CreateProduct {
	offer.Installments = installments
	return c
}

func (c *CreateProduct) SetOfferCoordinate(offer CreateOffers, coordinate entities.Coordinate) *CreateProduct {
	offer.Coordinate = coordinate
	return c
}

func (c *CreateProduct) Build() *CreateProduct {
	return c
}
