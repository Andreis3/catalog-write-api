//go:build unit

package entities_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/andreis3/catalog-write-api/internal/domain/entities"
)

var _ = Describe("INTERNAL :: DOMAIN :: ENTITIES :: PRODUCT", func() {
	Describe("#Validate", func() {
		Context("success cases", func() {
			It("should not return an error when valid entity", func() {
				entity := entities.ProductBuilder().
					SetID(1).
					SetApikeyID(1).
					SetExternalID("1").
					SetApikey("apikey").
					SetName("name").
					SetDescription("description").
					SetStatus("enabled").
					SetBrand("brand").
					SetReleaseDate("2021-01-01").
					Build()

				err := entity.Validate()

				Expect(err.Errors()).To(BeNil())
				Expect(err.HasErrors()).To(BeFalse())
				Expect(err.ListErrors()).To(BeEmpty())
			})
		})

		Context("error cases", func() {
			It("should return an error when invalid entity", func() {
				entity := entities.ProductBuilder().Build()

				err := entity.Validate()

				Expect(err.Errors()).To(Not(BeNil()))
				Expect(err.HasErrors()).To(BeTrue())
				Expect(err.ListErrors()).To(Not(BeEmpty()))
			})
		})
	})
})
