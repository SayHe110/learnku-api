package validator

import (
    "fmt"
    "github.com/go-playground/universal-translator"
    "gopkg.in/go-playground/validator.v9"
)

type CommonError struct {
    Errors map[string]interface{}
}

var (
    _isFiledExists = "该 %s 值不存在"
    _filedUnique   = "该 %s 值已存在"
)

func GetValidatorError(err error, trans ut.Translator) (res CommonError) {
    res = CommonError{}
    res.Errors = make(map[string]interface{})
    errs := err.(validator.ValidationErrors)

    for _, e := range errs {
        switch e.Tag() {
        case "filed_unique":
            res.Errors[e.Field()] = fmt.Sprintf(_filedUnique, e.Field())
            break
        case "is_filed_exist":
            res.Errors[e.Field()] = fmt.Sprintf(_isFiledExists, e.Field())
            break
        default:
            res.Errors[e.Field()] = e.Translate(trans)
        }
    }

    return
}
