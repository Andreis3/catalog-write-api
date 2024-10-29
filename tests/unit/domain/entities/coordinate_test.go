//go:build unit

package entities_test

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/andreis3/catalog-write-api/internal/domain/entities"
)

var _ = Describe("INTERNAL :: DOMAIN :: ENTITIES :: COORDINATE", func() {
	Describe("#Validate", func() {
		Context("success cases", func() {
			It("should not return an error when send all fields", func() {
				coordinate := entities.CoordinateBuilder().
					SetID(1).
					SetOffersID(1).
					SetLatitude(-29.69050937064539).
					SetLongitude(-51.23978484448051).
					Build()

				err := coordinate.Validate()

				Expect(err.Errors()).To(BeNil())
				Expect(err.HasErrors()).To(BeFalse())
				Expect(err.ListErrors()).To(BeEmpty())
			})
			It("should return method Get is called", func() {
				coordinate := entities.CoordinateBuilder().
					SetID(1).
					SetOffersID(1).
					SetLatitude(-29.69050937064539).
					SetLongitude(-51.23978484448051).
					Build()

				Expect(coordinate.GetID()).To(Equal(int64(1)))
				Expect(coordinate.GetOffersID()).To(Equal(int64(1)))
				Expect(coordinate.GetLatitude()).To(Equal(-29.69050937064539))
				Expect(coordinate.GetLongitude()).To(Equal(-51.23978484448051))
			})
		})

		Context("error cases", func() {
			It("should return an error when invalid latitude and longitude", func() {
				coordinate := entities.CoordinateBuilder().
					SetID(1).
					SetOffersID(1).
					SetLatitude(-180.1).
					SetLongitude(-200.1).
					Build()

				err := coordinate.Validate()

				Expect(err.Errors()).NotTo(BeNil())
				Expect(err.HasErrors()).To(BeTrue())
				Expect(err.Errors()).To(ContainElement(errors.New("latitude: must be between -90 and +90 degrees")))
				Expect(err.Errors()).To(ContainElement(errors.New("longitude: must be between -180 and +180 degrees")))
				Expect(err.ListErrors()).To(ContainSubstring("latitude: must be between -90 and +90 degrees"))
				Expect(err.ListErrors()).To(ContainSubstring("longitude: must be between -180 and +180 degrees"))
			})

			It("should return an error when latitude not set and longitude is set", func() {
				coordinate := entities.CoordinateBuilder().
					SetID(1).
					SetOffersID(1).
					SetLongitude(-51.23978484448051).
					Build()

				err := coordinate.Validate()

				Expect(err.Errors()).NotTo(BeNil())
				Expect(err.HasErrors()).To(BeTrue())
				Expect(err.Errors()).To(ContainElement(errors.New("latitude: must be set if longitude is set")))
				Expect(err.ListErrors()).To(ContainSubstring("latitude: must be set if longitude is set"))
			})

			It("should return an error when latitude is set and longitude not set", func() {
				coordinate := entities.CoordinateBuilder().
					SetID(1).
					SetOffersID(1).
					SetLatitude(-29.69050937064539).
					Build()

				err := coordinate.Validate()

				Expect(err.Errors()).NotTo(BeNil())
				Expect(err.HasErrors()).To(BeTrue())
				Expect(err.Errors()).To(ContainElement(errors.New("longitude: must be set if latitude is set")))
				Expect(err.ListErrors()).To(ContainSubstring("longitude: must be set if latitude is set"))
			})
		})
	})
})
