package util

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
)

type isEmptyRule string

const (
	NotAllowed isEmptyRule = ""
)

func (v isEmptyRule) Validate(value interface{}) error {
	if value == nil {
		return nil
	}

	value, isNil := validation.Indirect(value)
	if isNil || validation.IsEmpty(value) {
		return nil
	}

	return errors.Errorf("should be empty")
}

// func AppendErrors(errs validation.Errors, other validation.Errors) {
// 	for key, value := range other {
// 		errs[key] = value
// 	}
// }
