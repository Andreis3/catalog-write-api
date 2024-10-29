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
				installment := entities.InstallmentBuilder().
					SetID(1).
					SetOrderID(1).
					SetCount(1).
					SetPrice(1).
					Build()

				err := installment.Validate()

				Expect(err.Errors()).To(BeNil())
				Expect(err.HasErrors()).To(BeFalse())
				Expect(err.ListErrors()).To(BeEmpty())
			})

			It("should not return an error count and price equal 0", func() {
				installment := entities.InstallmentBuilder().
					SetID(1).
					SetOrderID(1).
					SetCount(0).
					SetPrice(0).
					Build()

				err := installment.Validate()

				Expect(err.Errors()).To(BeNil())
				Expect(err.HasErrors()).To(BeFalse())
				Expect(err.ListErrors()).To(BeEmpty())
			})
		})
		Context("error cases", func() {
			It("should return an error count and price less than 0", func() {
				installment := entities.InstallmentBuilder().
					SetID(1).
					SetOrderID(1).
					SetCount(-1).
					SetPrice(-1).
					Build()

				err := installment.Validate()

				Expect(err.Errors()).ToNot(BeNil())
				Expect(err.HasErrors()).To(BeTrue())
				Expect(err.ListErrors()).ToNot(BeEmpty())
				Expect(err.ListErrors()).To(ContainSubstring("price: price cannot be a negative number"))
				Expect(err.ListErrors()).To(ContainSubstring("count: count cannot be a negative number"))
				Expect(err.Errors()).To(ContainElement(errors.New("price: price cannot be a negative number")))
				Expect(err.Errors()).To(ContainElement(errors.New("count: count cannot be a negative number")))
			})

			It("should return an error when count is more than 12", func() {
				installment := entities.InstallmentBuilder().
					SetID(1).
					SetOrderID(1).
					SetCount(13).
					SetPrice(1).
					Build()

				err := installment.Validate()

				Expect(err.Errors()).ToNot(BeNil())
				Expect(err.HasErrors()).To(BeTrue())
				Expect(err.ListErrors()).ToNot(BeEmpty())
				Expect(err.ListErrors()).To(ContainSubstring("count: count cannot exceed 12"))
				Expect(err.Errors()).To(ContainElement(errors.New("count: count cannot exceed 12")))
			})
		})
	})
})