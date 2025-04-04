package sql

import (
	"database/sql"
	"time"
)

func NewNullInt64(value int64) sql.NullInt64 {
	return sql.NullInt64{
		Valid: true,
		Int64: value,
	}
}

func NewNullInt64Ptr(value *int64) sql.NullInt64 {
	var (
		result sql.NullInt64
	)

	if value != nil {
		result.Valid = true
		result.Int64 = *value
	} else {
		result.Valid = false
	}

	return result
}

func NewNullString(value string) sql.NullString {
	return sql.NullString{
		String: value,
		Valid:  true,
	}
}

func NewNullStringPtr(value *string) sql.NullString {
	var (
		result sql.NullString
	)

	if value != nil {
		result.Valid = true
		result.String = *value
	} else {
		result.Valid = false
	}

	return result
}

func NewNullBool(value bool) sql.NullBool {
	return sql.NullBool{
		Bool:  value,
		Valid: true,
	}
}

func NewNullTime(value time.Time) sql.NullTime {
	return sql.NullTime{
		Time:  value,
		Valid: true,
	}
}

func NewNullFloat64(value float64) sql.NullFloat64 {
	return sql.NullFloat64{
		Float64: value,
		Valid:   true,
	}
}

func NewNullFloat64Ptr(value *float64) sql.NullFloat64 {
	var (
		result sql.NullFloat64
	)

	if value != nil {
		result.Valid = true
		result.Float64 = *value
	} else {
		result.Valid = false
	}

	return result
}

func NewNullInt64Array(value []int64) []sql.NullInt64 {
	result := make([]sql.NullInt64, 0, len(value))

	for _, v := range value {
		result = append(result, NewNullInt64(v))
	}

	return result
}

func NewNullStringArray(value []string) []sql.NullString {
	result := make([]sql.NullString, 0, len(value))

	for _, v := range value {
		result = append(result, NewNullString(v))
	}

	return result
}
