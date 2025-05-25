package core_errors

import "fmt"

type DomainError struct {
	Code           int
	Message        string
	RussianMessage string
	Data           map[string]string
}

func (e *DomainError) Error() string {
	var (
		result string
	)

	if len(e.RussianMessage) > 0 {
		result = e.RussianMessage
	} else {
		result = e.Message
	}

	if e.Data != nil && len(e.Data) > 0 {
		dataValues := "."
		valuesInData := 0
		for key, value := range e.Data {
			if valuesInData > 0 {
				dataValues = dataValues + ","
			}
			dataValues += fmt.Sprintf(" %s: %s", key, value)
			valuesInData++
		}

		if valuesInData > 0 {
			result += dataValues
		}
	}

	return result
}

func NewDomainError(code int, message string, russianMessage string) *DomainError {
	result := &DomainError{
		Code:           code,
		Message:        message,
		RussianMessage: russianMessage,
	}

	return result
}

func (e *DomainError) AddValue(key, value string) *DomainError {
	if e.Data == nil {
		e.Data = make(map[string]string)
	}

	e.Data[key] = value

	return e
}
