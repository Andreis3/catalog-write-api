//go:build unit

package entities_test

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/andreis3/catalog-write-api/internal/domain/entities"
)

var _ = Describe("INTERNAL :: DOMAIN :: ENTITIES :: SPECIFICATION_KEY", func() {
	Describe("#Validate", func() {
		Context("success cases", func() {
			It("should not return an error when valid entity", func() {
				entity := entities.SpecificationKeyBuilder().
					SetID(1).
					SetKey("key").
					Build()

				err := entity.Validate()

				Expect(err.Errors()).To(BeNil())
				Expect(err.HasErrors()).To(BeFalse())
				Expect(err.ListErrors()).To(BeEmpty())
			})
		})

		Context("error cases", func() {
			It("should return an error when invalid entity", func() {
				entity := entities.SpecificationKeyBuilder().Build()

				err := entity.Validate()

				Expect(err.Errors()).To(Not(BeNil()))
				Expect(err.HasErrors()).To(BeTrue())
				Expect(err.ListErrors()).To(Not(BeEmpty()))
				Expect(err.ListErrors()).To(ContainSubstring("key: is required"))
				Expect(err.Errors()).To(HaveLen(1))
				Expect(err.Errors()).To(ContainElements([]error{
					errors.New("key: is required"),
				}))
			})

			It("should return an error when key contains more than 10 characters", func() {
				entity := entities.SpecificationKeyBuilder().
					SetKey("12345678901").
					Build()

				err := entity.Validate()

				Expect(err.Errors()).To(Not(BeNil()))
				Expect(err.HasErrors()).To(BeTrue())
				Expect(err.ListErrors()).To(Not(BeEmpty()))
				Expect(err.ListErrors()).To(ContainSubstring("key: limit max of the characters not more than 10"))
				Expect(err.Errors()).To(HaveLen(1))
				Expect(err.Errors()).To(ContainElements([]error{
					errors.New("key: limit max of the characters not more than 10"),
				}))
			})
		})
	})
})
