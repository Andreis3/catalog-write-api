//go:build unit

package entities_test

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/andreis3/catalog-write-api/internal/domain/entities"
)

var _ = Describe("INTERNAL :: DOMAIN :: ENTITIES :: INSTALLMENT", func() {
	Describe("#Validate", func() {
		Context("success cases", func() {
			It("should not return an error count and price more 0", func() {
				entity := entities.InstallmentBuilder().
					SetID(1).
					SetOrderID(1).
					SetCount(1).
					SetPrice(1).
					Build()

				err := entity.Validate()

				Expect(err.Errors()).To(BeNil())
				Expect(err.HasErrors()).To(BeFalse())
				Expect(err.ListErrors()).To(BeEmpty())
			})

			It("should not return an error count and price equal 0", func() {
				entity := entities.InstallmentBuilder().
					SetID(1).
					SetOrderID(1).
					SetCount(0).
					SetPrice(0).
					Build()

				err := entity.Validate()

				Expect(err.Errors()).To(BeNil())
				Expect(err.HasErrors()).To(BeFalse())
				Expect(err.ListErrors()).To(BeEmpty())
			})
		})
		Context("error cases", func() {
			It("should return an error count and price less than 0", func() {
				entity := entities.InstallmentBuilder().
					SetID(1).
					SetOrderID(1).
					SetCount(-1).
					SetPrice(-1).
					Build()

				err := entity.Validate()

				Expect(err.Errors()).ToNot(BeNil())
				Expect(err.HasErrors()).To(BeTrue())
				Expect(err.ListErrors()).ToNot(BeEmpty())
				Expect(err.ListErrors()).To(ContainSubstring("price: cannot be negative"))
				Expect(err.ListErrors()).To(ContainSubstring("count: cannot be negative"))
				Expect(err.Errors()).To(ContainElement(errors.New("price: cannot be negative")))
				Expect(err.Errors()).To(ContainElement(errors.New("count: cannot be negative")))
			})

			It("should return an error when count is more than 12", func() {
				entity := entities.InstallmentBuilder().
					SetID(1).
					SetOrderID(1).
					SetCount(13).
					SetPrice(1).
					Build()

				err := entity.Validate()

				Expect(err.Errors()).ToNot(BeNil())
				Expect(err.HasErrors()).To(BeTrue())
				Expect(err.ListErrors()).ToNot(BeEmpty())
				Expect(err.ListErrors()).To(ContainSubstring("count: cannot exceed 12"))
				Expect(err.Errors()).To(ContainElement(errors.New("count: cannot exceed 12")))
			})
		})
	})
})
