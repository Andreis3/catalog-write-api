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
		Context("when invalid apikey entity", func() {
			It("should return an error when apikey is empty", func() {
				apikey := entities.Builder().Build()

				err := apikey.Validate()

				Expect(err).NotTo(BeNil())
				Expect(err.Errors()).To(HaveLen(2))
				Expect(err.Errors()).To(ContainElement(errors.New("name: is required")))
				Expect(err.Errors()).To(ContainElement(errors.New("status: is required")))
				Expect(err.Error()).To(ContainSubstring("name: is required"))
				Expect(err.Error()).To(ContainSubstring("status: is required"))
			})

			It("should return an error when apikey status is invalid", func() {
				apikey := entities.Builder().SetID(1).SetName("apikey").SetStatus("invalid").Build()

				err := apikey.Validate()

				Expect(err).NotTo(BeNil())
				Expect(err.Errors()).To(HaveLen(1))
				Expect(err.Errors()).To(ContainElement(errors.New("status: is invalid, valid values are active or inactive")))
				Expect(err.Error()).To(ContainSubstring("status: is invalid, valid values are active or inactive"))
			})
		})
		Context("when valid apikey entity", func() {
			It("should not return an error when status is active", func() {
				apikey := entities.Builder().SetID(1).SetName("apikey").SetStatusActive().Build()

				err := apikey.Validate()

				Expect(err.Errors()).To(BeNil())
				Expect(err.HasErrors()).To(BeFalse())
				Expect(err.Error()).To(BeEmpty())
			})

			It("should not return an error when status is inactive", func() {
				apikey := entities.Builder().SetID(1).SetName("apikey").SetStatusInactive().Build()

				err := apikey.Validate()

				Expect(err.Errors()).To(BeNil())
				Expect(err.HasErrors()).To(BeFalse())
				Expect(err.Error()).To(BeEmpty())
			})
		})
	})
})
