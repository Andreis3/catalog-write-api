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
				Expect(err.Errors()).To(HaveLen(6))
				Expect(err.ListErrors()).To(ContainSubstring("external_id: is required"))
				Expect(err.ListErrors()).To(ContainSubstring("apikey: is required"))
				Expect(err.ListErrors()).To(ContainSubstring("name: is required"))
				Expect(err.ListErrors()).To(ContainSubstring("description: is required"))
				Expect(err.ListErrors()).To(ContainSubstring("status: is required"))
				Expect(err.ListErrors()).To(ContainSubstring("brand: is required"))
				Expect(err.Errors()).To(ContainElements([]error{
					errors.New("external_id: is required"),
					errors.New("apikey: is required"),
					errors.New("name: is required"),
					errors.New("description: is required"),
					errors.New("status: is required"),
					errors.New("brand: is required"),
				}))
			})

			It("should return an error when invalid status", func() {
				entity := entities.ProductBuilder().
					SetID(1).
					SetApikeyID(1).
					SetExternalID("1").
					SetApikey("apikey").
					SetName("name").
					SetDescription("description").
					SetStatus("invalid").
					SetBrand("brand").
					SetReleaseDate("2021-01-01").
					Build()

				err := entity.Validate()

				Expect(err.Errors()).To(Not(BeNil()))
				Expect(err.HasErrors()).To(BeTrue())
				Expect(err.ListErrors()).To(Not(BeEmpty()))
				Expect(err.Errors()).To(HaveLen(1))
				Expect(err.ListErrors()).To(ContainSubstring("status: is invalid, valid values are [enabled disabled]"))
				Expect(err.Errors()).To(ContainElements([]error{
					errors.New("status: is invalid, valid values are [enabled disabled]"),
				}))
			})

			It("should return an error when fields exceed max characters", func() {
				entity := entities.ProductBuilder().
					SetID(1).
					SetApikeyID(1).
					SetExternalID("1234564564165415648644846546" +
						"845614894984141449814914981491491" +
						"441894198149814981491491149814981494198419").
					SetApikey("apikey").
					SetName("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
						"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
						"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
						"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
						"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
						"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
						"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa").
					SetDescription("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
						"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
						"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
						"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
						"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
						"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
						"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa").
					SetStatus("enabled").
					SetBrand("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
						"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
						"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
						"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
						"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
						"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
						"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa").
					SetReleaseDate("2021-01-01").
					Build()

				err := entity.Validate()

				Expect(err.Errors()).To(Not(BeNil()))
				Expect(err.HasErrors()).To(BeTrue())
				Expect(err.ListErrors()).To(Not(BeEmpty()))
				Expect(err.Errors()).To(HaveLen(4))
				Expect(err.ListErrors()).To(ContainSubstring("external_id: limit max of the characters not more than 50"))
				Expect(err.ListErrors()).To(ContainSubstring("name: limit max of the characters not more than 50"))
				Expect(err.ListErrors()).To(ContainSubstring("description: limit max of the characters not more than 255"))
				Expect(err.ListErrors()).To(ContainSubstring("brand: limit max of the characters not more than 100"))
				Expect(err.Errors()).To(ContainElements([]error{
					errors.New("external_id: limit max of the characters not more than 50"),
					errors.New("name: limit max of the characters not more than 50"),
					errors.New("description: limit max of the characters not more than 255"),
					errors.New("brand: limit max of the characters not more than 100"),
				}))

			})
		})
	})
})
