package tools

import (
    "github.com/go-playground/universal-translator"
    "gopkg.in/go-playground/validator.v9"
)

type ValidatorCommonError struct {
    Errors map[string]interface{}
}

func GetValidatorError(err error, trans ut.Translator) (res ValidatorCommonError) {
    res = ValidatorCommonError{}
    res.Errors = make(map[string]interface{})
    errs := err.(validator.ValidationErrors)

    for _, e := range errs {
        res.Errors[e.Field()] = e.Translate(trans)
    }

    return
}
