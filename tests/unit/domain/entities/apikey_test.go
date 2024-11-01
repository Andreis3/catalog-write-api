//go:build unit

package entities_test

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/andreis3/catalog-write-api/internal/domain/entities"
)

var _ = Describe("INTERNAL :: DOMAIN :: ENTITIES :: APIKEY", func() {
	Describe("#Validate", func() {
		Context("success cases", func() {
			It("should not return an error when status is active", func() {
				entity := entities.ApiKeyBuilder().SetID(1).SetName("entity").Activate().Build()

				err := entity.Validate()

				Expect(err.Errors()).To(BeNil())
				Expect(err.HasErrors()).To(BeFalse())
				Expect(err.ListErrors()).To(BeEmpty())
			})

			It("should not return an error when status is inactive", func() {
				entity := entities.ApiKeyBuilder().SetID(1).SetName("apikey").Deactivate().Build()

				err := entity.Validate()

				Expect(err.Errors()).To(BeNil())
				Expect(err.HasErrors()).To(BeFalse())
				Expect(err.ListErrors()).To(BeEmpty())
			})

			It("should return method Get is called", func() {
				entity := entities.ApiKeyBuilder().SetID(1).SetName("apikey").Activate().Build()

				Expect(entity.GetID()).To(Equal(int64(1)))
				Expect(entity.GetName()).To(Equal("apikey"))
				Expect(entity.GetStatus()).To(Equal("active"))
			})
		})
		Context("error cases", func() {
			It("should return an error when apikey is empty", func() {
				apikey := entities.ApiKeyBuilder().Build()

				err := apikey.Validate()

				Expect(err).NotTo(BeNil())
				Expect(err.Errors()).To(HaveLen(2))
				Expect(err.Errors()).To(ContainElement(errors.New("name: is required")))
				Expect(err.Errors()).To(ContainElement(errors.New("status: is required")))
				Expect(err.ListErrors()).To(ContainSubstring("name: is required"))
				Expect(err.ListErrors()).To(ContainSubstring("status: is required"))
			})

			It("should return an error when apikey status is invalid", func() {
				apikey := entities.ApiKeyBuilder().SetID(1).SetName("apikey").SetStatus("invalid").Build()

				err := apikey.Validate()

				Expect(err).NotTo(BeNil())
				Expect(err.Errors()).To(HaveLen(1))
				Expect(err.Errors()).To(ContainElement(errors.New("status: is invalid, valid values are [active inactive]")))
				Expect(err.ListErrors()).To(ContainSubstring("status: is invalid, valid values are [active inactive]"))
			})

			It("should return an error when name is less than 3 characters", func() {
				apikey := entities.ApiKeyBuilder().SetID(1).SetName("a").Activate().Build()

				err := apikey.Validate()

				Expect(err).NotTo(BeNil())
				Expect(err.Errors()).To(HaveLen(1))
				Expect(err.Errors()).To(ContainElement(errors.New("name: limit min of the characters not less than 3")))
				Expect(err.ListErrors()).To(ContainSubstring("name: limit min of the characters not less than 3"))
			})
		})

	})
})
