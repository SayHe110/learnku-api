package validator

import "learnku-api/config"

func ValParams(params interface{}) (valError CommonError) {
    trans, _ := config.Uni.GetTranslator("zh")
    err := config.Validate.Struct(params)

    if err != nil {
        return GetValidatorError(err, trans)
    }

    return
}
