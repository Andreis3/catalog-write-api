//go:build unit

package entities_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/andreis3/catalog-write-api/internal/domain/entities"
)

var _ = Describe("INTERNAL :: DOMAIN :: ENTITIES :: OFFERS", func() {
	Describe("#Validate", func() {
		Context("success cases", func() {
			It("should not return an error when valid entity", func() {
				entity := entities.OfferBuilder().
					SetID(1).
					SetExternalID("external_id").
					SetSkuID(1).
					SetName("name").
					SetDescription("description").
					SetPrice(1.0).
					SetOldPrice(1.0).
					SetStock(1).
					SetStatus("available").
					SetSalesChannel("sales_channel").
					SetSeller("seller").
					SetStatusRemoved().
					SetStatusUnavailable().
					SetStatusAvailable().
					Build()

				err := entity.Validate()

				Expect(err.Errors()).To(BeNil())
				Expect(err.HasErrors()).To(BeFalse())
				Expect(err.ListErrors()).To(BeEmpty())
				Expect(entity.GetID()).To(Equal(int64(1)))
				Expect(entity.GetExternalID()).To(Equal("external_id"))
				Expect(entity.GetSkuID()).To(Equal(int64(1)))
				Expect(entity.GetName()).To(Equal("name"))
				Expect(entity.GetDescription()).To(Equal("description"))
				Expect(entity.GetPrice()).To(Equal(1.0))
				Expect(entity.GetOldPrice()).To(Equal(1.0))
				Expect(entity.GetStock()).To(Equal(int64(1)))
				Expect(entity.GetStatus()).To(Equal("available"))
				Expect(entity.GetSalesChannel()).To(Equal("sales_channel"))
				Expect(entity.GetSeller()).To(Equal("seller"))
			})
		})

		Context("error cases", func() {
			It("should return error when entity required fields are not set", func() {
				entity := entities.OfferBuilder().Build()

				err := entity.Validate()

				Expect(err.Errors()).ToNot(BeNil())
				Expect(err.HasErrors()).To(BeTrue())
				Expect(err.ListErrors()).ToNot(BeEmpty())
				Expect(err.ListErrors()).To(ContainSubstring("name: is required"))
				Expect(err.ListErrors()).To(ContainSubstring("external_id: is required"))
			})

			It("should return error when stock is negative", func() {
				entity := entities.OfferBuilder().
					SetID(1).
					SetExternalID("external_id").
					SetSkuID(1).
					SetName("name").
					SetDescription("description").
					SetPrice(1.0).
					SetOldPrice(1.0).
					SetStock(-1).
					SetStatus("available").
					SetSalesChannel("sales_channel").
					SetSeller("seller").
					Build()

				err := entity.Validate()

				Expect(err.Errors()).ToNot(BeNil())
				Expect(err.HasErrors()).To(BeTrue())
				Expect(err.ListErrors()).ToNot(BeEmpty())
				Expect(err.ListErrors()).To(ContainSubstring("stock: cannot be negative"))
			})

			It("should return error when price and old price are negative", func() {
				entity := entities.OfferBuilder().
					SetID(1).
					SetExternalID("external_id").
					SetSkuID(1).
					SetName("name").
					SetDescription("description").
					SetPrice(-1.0).
					SetOldPrice(-1.0).
					SetStock(1).
					SetStatus("available").
					SetSalesChannel("sales_channel").
					SetSeller("seller").
					Build()

				err := entity.Validate()

				Expect(err.Errors()).ToNot(BeNil())
				Expect(err.HasErrors()).To(BeTrue())
				Expect(err.ListErrors()).ToNot(BeEmpty())
				Expect(err.ListErrors()).To(ContainSubstring("price: cannot be negative"))
				Expect(err.ListErrors()).To(ContainSubstring("old_price: cannot be negative"))
			})

			It("should return error when name and description are too long", func() {
				entity := *entities.OfferBuilder().
					SetID(1).
					SetExternalID("external_id").
					SetSkuID(1).
					SetName("djjdiefkcemwnfknjioewjfoikewmjfnerogfvnreoikfjeroignfergvoerg" +
						"mfkewmfklemwflkermfgklrmeklgvmlkgmerklgfmkldmgvklmfrbvklmflkbgvermglkvmerger" +
						"dnkewfmdnklewmfklwenmflkewrmfklermfklwermfklewmfklmfklrmdlre").
					SetDescription("bdiebdfiuewnfiwnfduenbwdfbehfbdhfbhfboqbfheb fhd " +
						"bhvcb ewdhfbewfbdhebfhebwfhdefuijehfedufburhfure fuhuerferferbvfrbevfer" +
						"fnklwfnkldfmnlkwenmfknfjkernflkndklfcvvnkdjfkdsnvjknflkdvmfkldnvjksdfkvjnfdk" +
						"fnklwfnkldfmnlkwenmfknfjkernflkndklfcvvnkdjfkdsnvjknflkdvmfkldnvjksdfkvjnfdk").
					SetPrice(1.0).
					SetOldPrice(1.0).
					SetStock(1).
					SetStatus("available").
					SetSalesChannel("sales_channel").
					SetSeller("seller").
					Build()

				err := entity.Validate()

				Expect(err.Errors()).ToNot(BeNil())
				Expect(err.HasErrors()).To(BeTrue())
				Expect(err.ListErrors()).ToNot(BeEmpty())
				Expect(err.ListErrors()).To(ContainSubstring("name: limit max of the characters not more than 100"))
				Expect(err.ListErrors()).To(ContainSubstring("description: limit max of the characters not more than 255"))
			})

			It("should return error when status is invalid", func() {
				entity := entities.OfferBuilder().
					SetID(1).
					SetExternalID("external_id").
					SetSkuID(1).
					SetName("name").
					SetDescription("description").
					SetPrice(1.0).
					SetOldPrice(1.0).
					SetStock(1).
					SetStatus("invalid").
					SetSalesChannel("sales_channel").
					SetSeller("seller").
					Build()

				err := entity.Validate()

				Expect(err.Errors()).ToNot(BeNil())
				Expect(err.HasErrors()).To(BeTrue())
				Expect(err.ListErrors()).ToNot(BeEmpty())
				Expect(err.ListErrors()).To(ContainSubstring("is invalid, valid values are [available unavailable removed]"))
			})
		})
	})
})
