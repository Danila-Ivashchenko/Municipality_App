package validator

import "municipality_app/internal/common/validator/field"

type Validator struct {
	fields []field.Field
}

func (v *Validator) Validate() error {
	for _, f := range v.fields {
		err := f.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}

func (v *Validator) AddField(fields ...field.Field) {
	v.fields = append(v.fields, fields...)
}
