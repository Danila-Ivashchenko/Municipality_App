package field

import (
	"fmt"
	"municipality_app/internal/domain/core_errors"
)

type IntField struct {
	FieldTitle string
	Value      int

	checkers   []func() error
	IsRequired bool
}

func NewIntField(title string, value int) *IntField {
	return &IntField{
		FieldTitle: title,
		Value:      value,
	}
}

func (sf *IntField) Validate() error {
	for _, checker := range sf.checkers {
		if err := checker(); err != nil {
			return err
		}

		return nil
	}

	return nil
}

func (sf *IntField) Required() *IntField {
	checkFunc := func() error {
		if !(sf.Value != 0) {
			return core_errors.ValidationError.AddValue(sf.FieldTitle, fmt.Sprintf("обязательное поле"))
		}

		return nil
	}

	sf.checkers = append(sf.checkers, checkFunc)
	return sf
}

func (sf *IntField) Less(value int) *IntField {
	checkFunc := func() error {
		if !(sf.Value < value) {
			return core_errors.ValidationError.AddValue(sf.FieldTitle, fmt.Sprintf("значение должно быть меньше: %d", value))
		}

		return nil
	}

	sf.checkers = append(sf.checkers, checkFunc)
	return sf
}

func (sf *IntField) Bigger(value int) *IntField {
	checkFunc := func() error {
		if !(sf.Value > value) {
			return core_errors.ValidationError.AddValue(sf.FieldTitle, fmt.Sprintf("значение должно быть больше: %d", value))
		}

		return nil
	}

	sf.checkers = append(sf.checkers, checkFunc)
	return sf
}

func (sf *IntField) Equal(value int) *IntField {
	checkFunc := func() error {
		if !(sf.Value == value) {
			return core_errors.ValidationError.AddValue(sf.FieldTitle, fmt.Sprintf("значение должно быть равно: %d", value))
		}

		return nil
	}

	sf.checkers = append(sf.checkers, checkFunc)
	return sf
}

func (sf *IntField) Between(minValue, maxValue int) *IntField {
	checkFunc := func() error {
		if !(sf.Value >= minValue && sf.Value <= maxValue) {
			return core_errors.ValidationError.AddValue(sf.FieldTitle, fmt.Sprintf("значение должно быть между: %d и %d", minValue, maxValue))
		}

		return nil
	}

	sf.checkers = append(sf.checkers, checkFunc)
	return sf
}
