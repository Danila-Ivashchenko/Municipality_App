package field

import (
	"fmt"
	"municipality_app/internal/domain/core_errors"
)

type Int64Field struct {
	FieldTitle string
	Value      int64

	checkers   []func() error
	IsRequired bool
}

func NewInt64Field(title string, value int64) *Int64Field {
	return &Int64Field{
		FieldTitle: title,
		Value:      value,
	}
}

func (sf *Int64Field) Validate() error {
	for _, checker := range sf.checkers {
		if err := checker(); err != nil {
			return err
		}

		return nil
	}

	return nil
}

func (sf *Int64Field) Required() *Int64Field {
	checkFunc := func() error {
		if !(sf.Value != 0) {
			return core_errors.ValidationError.AddValue(sf.FieldTitle, fmt.Sprintf("обязательное поле"))
		}

		return nil
	}

	sf.checkers = append(sf.checkers, checkFunc)
	return sf
}

func (sf *Int64Field) Less(value int64) *Int64Field {
	checkFunc := func() error {
		if !(sf.Value < value) {
			return core_errors.ValidationError.AddValue(sf.FieldTitle, fmt.Sprintf("значение должно быть меньше: %d", value))
		}

		return nil
	}

	sf.checkers = append(sf.checkers, checkFunc)
	return sf
}

func (sf *Int64Field) Bigger(value int64) *Int64Field {
	checkFunc := func() error {
		if !(sf.Value > value) {
			return core_errors.ValidationError.AddValue(sf.FieldTitle, fmt.Sprintf("значение должно быть больше: %d", value))
		}

		return nil
	}

	sf.checkers = append(sf.checkers, checkFunc)
	return sf
}

func (sf *Int64Field) Equal(value int64) *Int64Field {
	checkFunc := func() error {
		if !(sf.Value == value) {
			return core_errors.ValidationError.AddValue(sf.FieldTitle, fmt.Sprintf("значение должно быть равно: %d", value))
		}

		return nil
	}

	sf.checkers = append(sf.checkers, checkFunc)
	return sf
}

func (sf *Int64Field) Between(minValue, maxValue int64) *Int64Field {
	checkFunc := func() error {
		if !(sf.Value >= minValue && sf.Value <= maxValue) {
			return core_errors.ValidationError.AddValue(sf.FieldTitle, fmt.Sprintf("значение должно быть между: %d и %d", minValue, maxValue))
		}

		return nil
	}

	sf.checkers = append(sf.checkers, checkFunc)
	return sf
}
