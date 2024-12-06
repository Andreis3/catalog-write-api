//go:build unit

package entities_test

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/andreis3/catalog-write-api/internal/domain/entities"
)

var _ = Describe("INTERNAL :: DOMAIN :: ENTITIES :: MEDIA", func() {
	Describe("#Validate", func() {
		Context("success cases", func() {
			It("should not return an error count and price more 0", func() {
				entity := entities.MediaBuilder().
					SetID(1).
					SetSkuID(1).
					SetURL("http://www.google.com").
					SetMediaType("image").
					SetDescription("description").
					SetIndex(1).
					SetMediaTypeVideo().
					SetMediaTypeImage().
					Build()

				err := entity.Validate()

				Expect(err.Errors()).To(BeNil())
				Expect(err.HasErrors()).To(BeFalse())
				Expect(err.ListErrors()).To(BeEmpty())
				Expect(entity.GetID()).To(Equal(int64(1)))
				Expect(entity.GetSkuID()).To(Equal(int64(1)))
				Expect(entity.GetURL()).To(Equal("http://www.google.com"))
				Expect(entity.GetMediaType()).To(Equal("image"))
				Expect(entity.GetDescription()).To(Equal("description"))
				Expect(entity.GetIndex()).To(Equal(1))
				Expect(entity.GetMediaType()).To(Equal("image"))
			})
		})

		Context("error cases", func() {
			It("should return error when entity is empty", func() {
				entity := entities.MediaBuilder().
					Build()

				err := entity.Validate()

				Expect(err.Errors()).ToNot(BeNil())
				Expect(err.HasErrors()).To(BeTrue())
				Expect(err.ListErrors()).ToNot(BeEmpty())
				Expect(err.ListErrors()).To(ContainSubstring("url: is required"))
				Expect(err.ListErrors()).To(ContainSubstring("media_type: is required"))
				Expect(err.ListErrors()).To(ContainSubstring("description: is required"))
				Expect(err.ListErrors()).To(ContainSubstring("index: cannot less than or equal to zero"))
				Expect(err.Errors()).To(ContainElement(errors.New("url: is required")))
				Expect(err.Errors()).To(ContainElement(errors.New("media_type: is required")))
				Expect(err.Errors()).To(ContainElement(errors.New("description: is required")))
				Expect(err.Errors()).To(ContainElement(errors.New("index: cannot less than or equal to zero")))
			})

			It("should return error when index equals 0", func() {
				entity := entities.MediaBuilder().
					SetID(1).
					SetSkuID(1).
					SetURL("http://www.google.com").
					SetMediaType("image").
					SetDescription("description").
					SetIndex(0).
					Build()

				err := entity.Validate()

				Expect(err.Errors()).To(ContainElement(errors.New("index: cannot less than or equal to zero")))
				Expect(err.HasErrors()).To(BeTrue())
				Expect(err.ListErrors()).To(ContainSubstring("index: cannot less than or equal to zero"))
			})
		})
	})
})
