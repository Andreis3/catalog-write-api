//go:build unit

package aggregates_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/andreis3/catalog-write-api/internal/domain/aggregates"
	"github.com/andreis3/catalog-write-api/internal/domain/entities"
)

var _ = Describe("INTERNAL :: DOMAIN :: AGGREGATES :: CREATE_PRODUCT", func() {
	Describe("#Validate", func() {
		Context("error cases", func() {
			It("should return an error when apikey is empty", func() {
				_, err := aggregates.CreateProductBuilder().
					WithProduct(entities.ProductBuilder().Build()).
					WithCategories([]*entities.Category{}).
					WithSkus([]aggregates.CreateSkus{}).
					Build()

				Expect(err).NotTo(BeNil())
				Expect(err.HasErrors()).To(BeTrue())
				Expect(err.Errors()).To(HaveLen(8))
				Expect(err.ListErrors()).To(ContainSubstring("product: external_id: is required"))
				Expect(err.ListErrors()).To(ContainSubstring("product: apikey: is required"))
				Expect(err.ListErrors()).To(ContainSubstring("product: name: is required"))
				Expect(err.ListErrors()).To(ContainSubstring("product: description: is required"))
				Expect(err.ListErrors()).To(ContainSubstring("product: brand: is required"))
				Expect(err.ListErrors()).To(ContainSubstring("product: status: is required"))
				Expect(err.ListErrors()).To(ContainSubstring("categories: minimum of 1 categories is required"))
				Expect(err.ListErrors()).To(ContainSubstring("skus: minimum of 1 skus is required"))
			})

			It("should return an error circular dependency", func() {
				_, err := aggregates.CreateProductBuilder().
					WithProduct(entities.ProductBuilder().
						SetApikey("apikey").
						SetExternalID("external_id").
						SetName("name").
						SetDescription("description").
						SetBrand("brand").
						SetStatus("enabled").
						Build()).
					WithCategories([]*entities.Category{
						entities.CategoryBuilder().SetID(1).SetCategoryKey("cat1").SetParentCategoryKey("").Build(),
						entities.CategoryBuilder().SetID(2).SetCategoryKey("cat2").SetParentCategoryKey("cat3").Build(),
						entities.CategoryBuilder().SetID(3).SetCategoryKey("cat3").SetParentCategoryKey("cat2").Build(),
					}).
					WithSkus([]aggregates.CreateSkus{
						{
							Skus: entities.SkuBuilder().SetID(1).SetExternalID("sku1").SetProductID(1).SetName("name").SetDescription("description").SetGtin("449849849").SetStatus("active").Build(),
							Medias: []*entities.Media{
								entities.MediaBuilder().SetID(1).SetSkuID(1).SetURL("http://url.com").SetMediaTypeImage().Build(),
							},
							Tags: []*entities.Tag{
								entities.TagBuilder().SetID(1).SetSkuID(1).Build(),
							},
							Specifications: []aggregates.CreateSpecifications{
								{
									SpecificationKey:   *entities.SpecificationKeyBuilder().SetID(1).SetAPIKeyID(1).SetKey("KEY").Build(),
									SpecificationValue: *entities.SpecificationValueBuilder().SetID(1).SetSpecificationKeyID(1).SetValue("VALUE").Build(),
								},
							},
						},
					}).
					Build()

				Expect(err).NotTo(BeNil())
				Expect(err.HasErrors()).To(BeTrue())
			})
		})
	})
})
