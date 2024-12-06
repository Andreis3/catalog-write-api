//go:build unit

package commons_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/andreis3/catalog-write-api/internal/domain/commons"
)

var _ = Describe("INTERNAL :: DOMAIN :: COMMONS :: VALIDATE_FIELDS", func() {
	Describe("#CheckIsValidStatus", func() {
		Context("error cases", func() {
			It("should return an error", func() {
				validateFields := commons.NewValidateFields()
				err := validateFields.CheckIsValidStatus("status", "name", []string{"any", "any2", "any3"})

				Expect(err).NotTo(BeNil())
				Expect(err).To(MatchError("name: is invalid, valid values are [any any2 any3]"))
			})
		})
	})

	Describe("#CheckEmptyField", func() {
		Context("error cases", func() {
			It("should return an error", func() {
				validateFields := commons.NewValidateFields()
				err := validateFields.CheckEmptyField("", "name")

				Expect(err).NotTo(BeNil())
				Expect(err).To(MatchError("name: is required"))
			})
		})
	})

	Describe("#CheckMaxCharacters", func() {
		Context("error cases", func() {
			It("should return an error", func() {
				validateFields := commons.NewValidateFields()
				err := validateFields.CheckMaxCharacters("field", "name", 3)

				Expect(err).NotTo(BeNil())
				Expect(err).To(MatchError("name: limit max of the characters not more than 3"))
			})
		})
	})

	Describe("#CheckMinCharacters", func() {
		Context("error cases", func() {
			It("should return an error", func() {
				validateFields := commons.NewValidateFields()
				err := validateFields.CheckMinCharacters("field", "name", 10)

				Expect(err).NotTo(BeNil())
				Expect(err).To(MatchError("name: limit min of the characters not less than 10"))
			})
		})
	})

	Describe("#CheckLatitudeRange", func() {
		Context("error cases", func() {
			It("should return an error", func() {
				validateFields := commons.NewValidateFields()
				err := validateFields.CheckLatitudeRange(91.0)

				Expect(err).NotTo(BeNil())
				Expect(err).To(MatchError("latitude: must be between -90 and +90 degrees"))
			})
		})
	})

	Describe("#CheckLongitudeRange", func() {
		Context("error cases", func() {
			It("should return an error", func() {
				validateFields := commons.NewValidateFields()
				err := validateFields.CheckLongitudeRange(181.0)

				Expect(err).NotTo(BeNil())
				Expect(err).To(MatchError("longitude: must be between -180 and +180 degrees"))
			})
		})
	})

	Describe("#CheckMinimumOfOne", func() {
		Context("error cases", func() {
			It("should return an error", func() {
				validateFields := commons.NewValidateFields()
				err := validateFields.CheckMinimumOfOne(0, "name")

				Expect(err).NotTo(BeNil())
				Expect(err).To(MatchError("name: minimum of 1 name is required"))
			})
		})

		Context("success cases", func() {
			It("should return nil", func() {
				validateFields := commons.NewValidateFields()
				err := validateFields.CheckMinimumOfOne(1, "name")

				Expect(err).To(BeNil())
			})
		})
	})

	Describe("#CheckFieldEqualZero", func() {
		Context("error cases", func() {
			It("should return an error int", func() {
				validateFields := commons.NewValidateFields()
				err := validateFields.CheckFieldEqualZero(0, "name")

				Expect(err).NotTo(BeNil())
				Expect(err).To(MatchError("name: cannot less than or equal to zero"))
			})

			It("should return an error float64", func() {
				validateFields := commons.NewValidateFields()
				err := validateFields.CheckFieldEqualZero(0.0, "name")

				Expect(err).NotTo(BeNil())
				Expect(err).To(MatchError("name: cannot less than or equal to zero"))
			})

			It("should return an error int64", func() {
				validateFields := commons.NewValidateFields()
				err := validateFields.CheckFieldEqualZero(int64(0), "name")

				Expect(err).NotTo(BeNil())
				Expect(err).To(MatchError("name: cannot less than or equal to zero"))
			})

			Context("success cases", func() {
				It("should return nil", func() {
					validateFields := commons.NewValidateFields()
					err := validateFields.CheckFieldEqualZero(1, "name")

					Expect(err).To(BeNil())
				})
			})
		})

		Describe("#CheckNegativeField", func() {
			Context("error cases", func() {
				It("should return an error", func() {
					validateFields := commons.NewValidateFields()
					err := validateFields.CheckNegativeField(-1, "name")

					Expect(err).NotTo(BeNil())
					Expect(err).To(MatchError("name: cannot be negative"))
				})
			})

			Context("success cases", func() {
				It("should return nil", func() {
					validateFields := commons.NewValidateFields()
					err := validateFields.CheckNegativeField(1, "name")

					Expect(err).To(BeNil())
				})
			})
		})
	})

	Describe("#CheckSetField", func() {
		Context("error cases", func() {
			It("should return an error is string", func() {
				validateFields := commons.NewValidateFields()
				fieldNotSet := ""
				fieldSet := "fieldSet"
				name := "name"
				nameSet := "nameSet"
				err := validateFields.CheckSetField(fieldNotSet, fieldSet, name, nameSet)

				Expect(err).NotTo(BeNil())
				Expect(err).To(MatchError("name: must be set if nameSet is set"))
			})

			It("should return an error is float64", func() {
				validateFields := commons.NewValidateFields()
				fieldNotSet := 0.0
				fieldSet := 1.0
				name := "name"
				nameSet := "nameSet"
				err := validateFields.CheckSetField(fieldNotSet, fieldSet, name, nameSet)

				Expect(err).NotTo(BeNil())
				Expect(err).To(MatchError("name: must be set if nameSet is set"))
			})
		})

		Context("success cases", func() {
			It("should return nil is string", func() {
				validateFields := commons.NewValidateFields()
				fieldNotSet := "fieldNotSet"
				fieldSet := "fieldSet"
				name := "name"
				nameSet := "nameSet"
				err := validateFields.CheckSetField(fieldNotSet, fieldSet, name, nameSet)

				Expect(err).To(BeNil())
			})

			It("should return nil is float64", func() {
				validateFields := commons.NewValidateFields()
				fieldNotSet := 1.0
				fieldSet := 1.0
				name := "name"
				nameSet := "nameSet"
				err := validateFields.CheckSetField(fieldNotSet, fieldSet, name, nameSet)

				Expect(err).To(BeNil())
			})
		})
	})

	Describe("#CheckCircularDependencies", func() {
		Context("error cases", func() {
			It("should return an error", func() {
				validateFields := commons.NewValidateFields()
				err := validateFields.CheckCircularDependencies("name")

				Expect(err).NotTo(BeNil())
				Expect(err).To(MatchError("name: invalid circular dependencies"))
			})
		})
	})
})
