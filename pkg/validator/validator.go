package validator

import "learnku-api/config"

// 该方法为统一参数验证，若有自定义验证方法，那么该方法不可用，需单独写
func ValParams(params interface{}) (valError CommonError) {
    trans, _ := config.Uni.GetTranslator("zh")
    err := config.Validate.Struct(params)

    if err != nil {
        return GetValidatorError(err, trans)
    }

    return
}
