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
			})
		})
	})
})
