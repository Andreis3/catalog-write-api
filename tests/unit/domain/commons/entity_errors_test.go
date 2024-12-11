//go:build unit

package commons_test

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/andreis3/catalog-write-api/internal/domain/commons"
)

var _ = Describe("INTERNAL :: DOMAIN :: COMMONS :: ENTITY_ERRORS", func() {
	Describe("#Add, #HasErrors, #Errors, #ListErrors", func() {
		Context("success cases", func() {
			It("should return nil", func() {
				entityErrors := commons.NewEntityErrors()
				entityErrors.Add(nil)

				Expect(entityErrors.HasErrors()).To(BeFalse())
				Expect(entityErrors.Errors()).To(HaveLen(0))
			})
		})

		Context("error cases", func() {
			It("should add a error", func() {
				entityErrors := commons.NewEntityErrors()
				entityErrors.Add(errors.New("name: is invalid, valid values are [any any2 any3]"))

				Expect(entityErrors.HasErrors()).To(BeTrue())
				Expect(entityErrors.Errors()).To(HaveLen(1))
				Expect(entityErrors.ListErrors()).To(Equal("name: is invalid, valid values are [any any2 any3]"))
			})
		})
	})

	Describe("#MergeSlice", func() {
		Context("success cases", func() {
			It("should merge a slice of errors", func() {
				entityErrors := commons.NewEntityErrors()
				childErrors := commons.NewEntityErrors()
				childErrors.Add(errors.New("is invalid, valid values are [any any2 any3]"))

				entityErrors.MergeSlice(0, "name", childErrors)

				Expect(entityErrors.HasErrors()).To(BeTrue())
				Expect(entityErrors.Errors()).To(HaveLen(1))
				Expect(entityErrors.ListErrors()).To(Equal("name[0].name: is invalid, valid values are [any any2 any3]"))
			})
		})
	})
})
