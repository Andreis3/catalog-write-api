//go:build unit

package entities_test

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/andreis3/catalog-write-api/internal/domain/entities"
)

var _ = Describe("INTERNAL :: DOMAIN :: ENTITIES :: SPECIFICATION_VALUE", func() {
	Describe("#Validate", func() {
		Context("success cases", func() {
			It("should not return an error when valid entity", func() {
				entity := entities.SpecificationValueBuilder().
					SetID(1).
					SetSpecificationKeyID(1).
					SetValue("value").
					Build()

				err := entity.Validate()

				Expect(err.Errors()).To(BeNil())
				Expect(err.HasErrors()).To(BeFalse())
				Expect(err.ListErrors()).To(BeEmpty())
				Expect(entity.GetID()).To(Equal(int64(1)))
				Expect(entity.GetSpecificationKeyID()).To(Equal(int64(1)))
				Expect(entity.GetValue()).To(Equal("value"))
			})
		})

		Context("error cases", func() {
			It("should return an error when invalid entity", func() {
				entity := entities.SpecificationValueBuilder().Build()

				err := entity.Validate()

				Expect(err.Errors()).To(Not(BeNil()))
				Expect(err.HasErrors()).To(BeTrue())
				Expect(err.ListErrors()).To(Not(BeEmpty()))
				Expect(err.ListErrors()).To(ContainSubstring("value: is required"))
				Expect(err.Errors()).To(HaveLen(1))
				Expect(err.Errors()).To(ContainElements([]error{
					errors.New("value: is required"),
				}))
			})

			It("should return an error when key contains more than 10 characters", func() {
				entity := entities.SpecificationValueBuilder().
					SetValue("12345678901").
					Build()

				err := entity.Validate()

				Expect(err.Errors()).To(Not(BeNil()))
				Expect(err.HasErrors()).To(BeTrue())
				Expect(err.ListErrors()).To(Not(BeEmpty()))
				Expect(err.ListErrors()).To(ContainSubstring("value: limit max of the characters not more than 10"))
				Expect(err.Errors()).To(HaveLen(1))
				Expect(err.Errors()).To(ContainElements([]error{
					errors.New("value: limit max of the characters not more than 10"),
				}))
			})
		})
	})
})
