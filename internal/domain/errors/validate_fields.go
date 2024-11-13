package errors

import (
	"fmt"
	"slices"
)

const (
	STATUS_INVALID                 = "is invalid, valid values are"
	IS_REQUIRED                    = "is required"
	LIMIT_CHARACTERS_NOT_MORE_THAN = "limit max of the characters not more than"
	LIMIT_CHARACTERS_NOT_LESS_THAN = "limit min of the characters not less than"
	LATITUDE_OUT_OF_RANGE          = "must be between -90 and +90 degrees"
	LONGITUDE_OUT_OF_RANGE         = "must be between -180 and +180 degrees"
	SET_FIELD_NOT_SET_OUTHER_FIELD = "must be set if %s is set"
	CANNOT_NEGATIVE                = "cannot be negative"
	CANNOT_EXCEED                  = "cannot exceed"
	CANNOT_EQUAL_ZERO              = "cannot less than or equal to zero"
	MINIMUM_OF_ONE                 = "minimum of 1 %s is required"
	CIRCULAR_DEPENDENCIES          = "invalid circular dependencies"
)

type ValidateFields struct{}

func (v *ValidateFields) CheckIsValidStatus(field, name string, status []string) error {
	if field != "" && !slices.Contains(status, field) {
		return fmt.Errorf("%s: %s %s", name, STATUS_INVALID, status)
	}
	return nil
}

func (v *ValidateFields) CheckEmptyField(field string, name string) error {
	if field == "" {
		return fmt.Errorf("%s: %s", name, IS_REQUIRED)
	}
	return nil
}

func (v *ValidateFields) CheckMaxCharacters(field, name string, limit int) error {
	if len(field) > limit {
		return fmt.Errorf("%s: %s %d", name, LIMIT_CHARACTERS_NOT_MORE_THAN, limit)
	}
	return nil
}

func (v *ValidateFields) CheckMinCharacters(field, name string, limit int) error {
	if field != "" && len(field) < limit {
		return fmt.Errorf("%s: %s %d", name, LIMIT_CHARACTERS_NOT_LESS_THAN, limit)
	}
	return nil
}

func (v *ValidateFields) CheckSetField(fieldNotSet, fieldSet any, name, nameSet string) error {
	switch {
	case v.isString(fieldSet) && v.isString(fieldNotSet):
		if fieldSet.(string) != "" && fieldNotSet.(string) == "" {
			setField := fmt.Sprintf(SET_FIELD_NOT_SET_OUTHER_FIELD, nameSet)
			return fmt.Errorf("%s: %s", name, setField)
		}
	case v.isFloat64(fieldSet) && v.isFloat64(fieldNotSet):
		if fieldSet.(float64) != 0 && fieldNotSet.(float64) == 0 {
			setField := fmt.Sprintf(SET_FIELD_NOT_SET_OUTHER_FIELD, nameSet)
			return fmt.Errorf("%s: %s", name, setField)
		}
	}

	return nil
}

func (v *ValidateFields) CheckLatitudeRange(latitude float64) error {
	if latitude != 0 && (latitude < -90 || latitude > 90) {
		return fmt.Errorf("latitude: %s", LATITUDE_OUT_OF_RANGE)
	}
	return nil
}

func (v *ValidateFields) CheckLongitudeRange(longitude float64) error {
	if longitude != 0 && (longitude < -180 || longitude > 180) {
		return fmt.Errorf("longitude: %s", LONGITUDE_OUT_OF_RANGE)
	}
	return nil
}

func (v *ValidateFields) CheckNegativeField(field any, name string) error {
	switch field.(type) {
	case int64:
		if field.(int64) < 0 {
			return fmt.Errorf("%s: %s", name, CANNOT_NEGATIVE)
		}
	case float64:
		if field.(float64) < 0 {
			return fmt.Errorf("%s: %s", name, CANNOT_NEGATIVE)
		}
	case int:
		if field.(int) < 0 {
			return fmt.Errorf("%s: %s", name, CANNOT_NEGATIVE)
		}
	default:
		return nil
	}
	return nil
}

func (v *ValidateFields) CheckExceedField(field int, name string, limit int) error {
	if field > limit {
		return fmt.Errorf("%s: %s %d", name, CANNOT_EXCEED, limit)
	}
	return nil
}

func (v *ValidateFields) CheckMinimumOfOne(length int, name string) error {
	if length == 0 {
		return fmt.Errorf("%s: %s", name, fmt.Sprintf(MINIMUM_OF_ONE, name))
	}
	return nil
}

func (v *ValidateFields) CheckFieldEqualZero(field any, name string) error {
	switch field.(type) {
	case int64:
		if field.(int64) == 0 {
			return fmt.Errorf("%s: %s", name, CANNOT_EQUAL_ZERO)
		}
	case float64:
		if field.(float64) == 0 {
			return fmt.Errorf("%s: %s", name, CANNOT_EQUAL_ZERO)
		}
	case int:
		if field.(int) == 0 {
			return fmt.Errorf("%s: %s", name, CANNOT_EQUAL_ZERO)
		}
	default:
		return nil
	}
	return nil
}

func (v *ValidateFields) isString(field any) bool {
	_, ok := field.(string)
	return ok
}

func (v *ValidateFields) isFloat64(field any) bool {
	_, ok := field.(float64)
	return ok
}

func (v *ValidateFields) CheckCircularDependencies(s string) error {
	return fmt.Errorf("%s: %s", s, CIRCULAR_DEPENDENCIES)
}
