//go:build unit

package entities_test

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/andreis3/catalog-write-api/internal/domain/entities"
)

var _ = Describe("INTERNAL :: DOMAIN :: ENTITIES :: CATEGORY", func() {
	Describe("#Validate", func() {
		Context("success cases", func() {
			It("should not return an error when send all fields", func() {
				entity := entities.CategoryBuilder().
					SetID(1).
					SetCategoryKey("categoryKey").
					SetDescription("description").
					SetParentID(1).
					SetParentCategoryKey("parentCategoryKey").
					Build()

				err := entity.Validate()

				Expect(err.Errors()).To(BeNil())
				Expect(err.HasErrors()).To(BeFalse())
				Expect(err.ListErrors()).To(BeEmpty())
			})

			It("should not return an error when send only required fields", func() {
				entity := entities.CategoryBuilder().
					SetCategoryKey("categoryKey").
					SetDescription("description").
					Build()

				err := entity.Validate()

				Expect(err.Errors()).To(BeNil())
				Expect(err.HasErrors()).To(BeFalse())
				Expect(err.ListErrors()).To(BeEmpty())
			})

			It("should return method Get is called", func() {
				entity := entities.CategoryBuilder().
					SetID(1).
					SetCategoryKey("categoryKey").
					SetAPIKeyID(1).
					SetDescription("description").
					SetParentID(1).
					SetParentCategoryKey("parentCategoryKey").
					Build()

				Expect(entity.GetID()).To(Equal(int64(1)))
				Expect(entity.GetAPIKeyID()).To(Equal(int64(1)))
				Expect(entity.GetCategoryKey()).To(Equal("categoryKey"))
				Expect(entity.GetDescription()).To(Equal("description"))
				Expect(entity.GetParentID()).To(Equal(int64(1)))
				Expect(entity.GetParentCategoryKey()).To(Equal("parentCategoryKey"))
			})
		})
		Context("error cases", func() {
			It("should return an error when not send required fields", func() {
				entity := entities.CategoryBuilder().
					SetID(1).
					SetParentID(1).
					SetParentCategoryKey("parentCategoryKey").
					Build()

				err := entity.Validate()

				Expect(err).NotTo(BeNil())
				Expect(err.Errors()).To(HaveLen(2))
				Expect(err.Errors()).To(ContainElement(errors.New("category_key: is required")))
				Expect(err.Errors()).To(ContainElement(errors.New("description: is required")))
				Expect(err.ListErrors()).To(ContainSubstring("category_key: is required"))
				Expect(err.ListErrors()).To(ContainSubstring("description: is required"))
			})
		})
	})
})
