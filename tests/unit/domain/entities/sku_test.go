//go:build unit

package entities_test

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/andreis3/catalog-write-api/internal/domain/entities"
)

var _ = Describe("INTERNAL :: DOMAIN :: ENTITIES :: PRODUCT", func() {
	Describe("#Validate", func() {
		Context("success cases", func() {
			It("should not return an error when valid entity", func() {
				entity := entities.SkuBuilder().
					SetID(1).
					SetExternalID("external-id").
					SetProductID(1).
					SetName("name").
					SetDescription("description").
					SetGtin("gtin").
					SetStatus("active").
					SetStatusInactive().
					SetStatusActive().
					Build()

				err := entity.Validate()

				Expect(err.Errors()).To(BeNil())
				Expect(err.HasErrors()).To(BeFalse())
				Expect(err.ListErrors()).To(BeEmpty())
				Expect(entity.GetID()).To(Equal(int64(1)))
				Expect(entity.GetExternalID()).To(Equal("external-id"))
				Expect(entity.GetProductID()).To(Equal(int64(1)))
				Expect(entity.GetName()).To(Equal("name"))
				Expect(entity.GetDescription()).To(Equal("description"))
				Expect(entity.GetGtin()).To(Equal("gtin"))
				Expect(entity.GetStatus()).To(Equal("active"))
			})
		})

		Context("error cases", func() {
			It("should return an error when invalid entity", func() {
				entity := entities.SkuBuilder().
					SetID(1).
					SetExternalID("").
					SetProductID(1).
					SetName("").
					SetDescription("").
					SetGtin("").
					SetStatus("").
					Build()

				err := entity.Validate()

				Expect(err.Errors()).ToNot(BeNil())
				Expect(err.HasErrors()).To(BeTrue())
				Expect(err.ListErrors()).ToNot(BeEmpty())
				Expect(err.ListErrors()).To(ContainSubstring("external_id: is required"))
				Expect(err.ListErrors()).To(ContainSubstring("name: is required"))
				Expect(err.ListErrors()).To(ContainSubstring("description: is required"))
				Expect(err.ListErrors()).To(ContainSubstring("gtin: is required"))
				Expect(err.ListErrors()).To(ContainSubstring("status: is required"))
				Expect(err.Errors()).To(HaveLen(5))
				Expect(err.Errors()).To(ContainElements([]error{
					errors.New("external_id: is required"),
					errors.New("name: is required"),
					errors.New("description: is required"),
					errors.New("gtin: is required"),
					errors.New("status: is required"),
				}))
			})

			It("should return an error when invalid status", func() {
				entity := entities.SkuBuilder().
					SetID(1).
					SetExternalID("external-id").
					SetProductID(1).
					SetName("name").
					SetDescription("description").
					SetGtin("gtin").
					SetStatus("invalid").
					Build()

				err := entity.Validate()

				Expect(err.Errors()).ToNot(BeNil())
				Expect(err.HasErrors()).To(BeTrue())
				Expect(err.ListErrors()).ToNot(BeEmpty())
				Expect(err.Errors()).To(HaveLen(1))
				Expect(err.ListErrors()).To(ContainSubstring("status: is invalid, valid values are [active inactive]"))
				Expect(err.Errors()).To(ContainElement(errors.New("status: is invalid, valid values are [active inactive]")))
			})

			It("should return an error when fields exceed max characters", func() {
				entity := entities.SkuBuilder().
					SetID(1).
					SetExternalID("123456789012345678901234567890123456789012345678901").
					SetProductID(1).
					SetName("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
						"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
						"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
						"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
						"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
						"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
						"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa").
					SetDescription("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
						"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
						"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
						"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
						"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
						"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
						"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa").
					SetGtin("gtin").
					SetStatus("active").
					Build()

				err := entity.Validate()

				Expect(err.Errors()).ToNot(BeNil())
				Expect(err.HasErrors()).To(BeTrue())
				Expect(err.ListErrors()).ToNot(BeEmpty())
				Expect(err.Errors()).To(HaveLen(3))
				Expect(err.ListErrors()).To(ContainSubstring("external_id: limit max of the characters not more than 20"))
				Expect(err.ListErrors()).To(ContainSubstring("name: limit max of the characters not more than 100"))
				Expect(err.ListErrors()).To(ContainSubstring("description: limit max of the characters not more than 255"))
				Expect(err.Errors()).To(ContainElements([]error{
					errors.New("external_id: limit max of the characters not more than 20"),
					errors.New("name: limit max of the characters not more than 100"),
					errors.New("description: limit max of the characters not more than 255"),
				}))
			})
		})
	})
})
