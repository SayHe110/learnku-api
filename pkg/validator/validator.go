package validator

import (
    "github.com/jinzhu/gorm"
    "gopkg.in/go-playground/validator.v9"
    "learnku-api/config"
    "learnku-api/pkg/db"
    "strings"
)

var (
    tableNameVar string
    conn         *gorm.DB
    count        int
)

// 该方法为统一参数验证，若有自定义验证方法，那么该方法不可用，需单独写
func ValParams(params interface{}) (valError CommonError) {
    trans, _ := config.C.Uni.GetTranslator("zh")
    err := config.C.Validate.Struct(params)

    if err != nil {
        return GetValidatorError(err, trans)
    }

    return
}

func CustomValParams(params interface{}, customFunc validator.Func, tagName, tableName string) (valError CommonError) {
    tableNameVar, conn, count = tableName, db.NewMysql(config.C), 0 // 这样赋值不知道会不会挨打 ╮(╯▽╰)╭

    trans, _ := config.C.Uni.GetTranslator("zh")
    if regErr := config.C.Validate.RegisterValidation(tagName, customFunc); regErr != nil {
        panic(regErr)
    }

    err := config.C.Validate.Struct(params)

    if err != nil {
        return GetValidatorError(err, trans)
    }

    return
}

//TODO 为啥会执行两遍
func FiledValUnique(fl validator.FieldLevel) bool {
    value := fl.Field().String()
    column := strings.ToLower(fl.FieldName())

    conn.Table(tableNameVar).Where(column+"= ? ", value).Count(&count)

    return ! (count > 0)
}

func FiledValExists(fl validator.FieldLevel) bool {
    value := fl.Field().String()
    column := strings.ToLower(fl.FieldName())

    conn.Table(tableNameVar).Where(column+"= ? ", value).Count(&count)

    return count > 0
}
