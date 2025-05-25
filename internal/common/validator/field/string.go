package field

import (
	"fmt"
	"municipality_app/internal/domain/core_errors"
)

type StringField struct {
	FieldTitle string
	Value      string

	checkers   []func() error
	IsRequired bool
}

func NewStringField(title string, value string) *StringField {
	return &StringField{
		FieldTitle: title,
		Value:      value,
	}
}

func (sf *StringField) Validate() error {
	for _, checker := range sf.checkers {
		if err := checker(); err != nil {
			return err
		}

		return nil
	}

	return nil
}

func (sf *StringField) Required() *StringField {
	checkFunc := func() error {
		if !(len(sf.Value) != 0) {
			return core_errors.ValidationError.AddValue(sf.FieldTitle, fmt.Sprintf("обязательный параметр"))
		}

		return nil
	}

	sf.checkers = append(sf.checkers, checkFunc)
	return sf
}

func (sf *StringField) Less(length int) *StringField {
	checkFunc := func() error {
		if !(len(sf.Value) < length) {
			return core_errors.ValidationError.AddValue(sf.FieldTitle, fmt.Sprintf("длина должна быть меньше: %d", length))
		}

		return nil
	}

	sf.checkers = append(sf.checkers, checkFunc)
	return sf
}

func (sf *StringField) Bigger(length int) *StringField {
	checkFunc := func() error {
		if !(len(sf.Value) > length) {
			return core_errors.ValidationError.AddValue(sf.FieldTitle, fmt.Sprintf("длина должна быть больше: %d", length))
		}

		return nil
	}

	sf.checkers = append(sf.checkers, checkFunc)
	return sf
}

func (sf *StringField) Equal(length int) *StringField {
	checkFunc := func() error {
		if !(len(sf.Value) == length) {
			return core_errors.ValidationError.AddValue(sf.FieldTitle, fmt.Sprintf("длина должна быть равна: %d", length))
		}

		return nil
	}

	sf.checkers = append(sf.checkers, checkFunc)
	return sf
}

func (sf *StringField) Between(minLength, maxLength int) *StringField {
	checkFunc := func() error {
		if !(len(sf.Value) >= minLength && len(sf.Value) <= maxLength) {
			return core_errors.ValidationError.AddValue(sf.FieldTitle, fmt.Sprintf("длина должна быть между: %d и %d", minLength, maxLength))
		}

		return nil
	}

	sf.checkers = append(sf.checkers, checkFunc)
	return sf
}
