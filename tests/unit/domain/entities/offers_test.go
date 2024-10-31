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
					Build()

				err := entity.Validate()

				Expect(err.Errors()).To(BeNil())
				Expect(err.HasErrors()).To(BeFalse())
				Expect(err.ListErrors()).To(BeEmpty())
			})
		})

		Context("error cases", func() {
			It("should return error when entity required fields are not set", func() {
				entity := entities.OfferBuilder().Build()

				err := entity.Validate()

				Expect(err.Errors()).ToNot(BeNil())
				Expect(err.HasErrors()).To(BeTrue())
				Expect(err.ListErrors()).ToNot(BeEmpty())
				Expect(err.ListErrors()).To(ContainSubstring("name: Name is required"))
				Expect(err.ListErrors()).To(ContainSubstring("external_id: ExternalID is required"))
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
				Expect(err.ListErrors()).To(ContainSubstring("stock: Stock must be greater than or equal to 0"))
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
				Expect(err.ListErrors()).To(ContainSubstring("price: Price is required"))
				Expect(err.ListErrors()).To(ContainSubstring("old_price: Old price is required"))
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
				Expect(err.ListErrors()).To(ContainSubstring("name: Name must have a maximum of 100 characters"))
				Expect(err.ListErrors()).To(ContainSubstring("description: Description must have a maximum of 255 characters"))
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
				Expect(err.ListErrors()).To(ContainSubstring("status: Invalid status"))
			})
		})
	})
})
