package config

import (
    "github.com/go-playground/locales/zh"
    "github.com/go-playground/universal-translator"
    "github.com/spf13/viper"
    "gopkg.in/go-playground/validator.v9"
    zhTranslations "gopkg.in/go-playground/validator.v9/translations/zh"
    "log"
)

const (
    CONFIGFILENAME = "./.env.yaml"
    CONFIGFILETYPE = "yaml"
)

var (
    AppConfig *appConfig
    DBConfig  *dbConfig
    Uni       *ut.UniversalTranslator
    Validate  *validator.Validate
)

func Init(filename string) {
    if filename == "" {
        filename = CONFIGFILENAME
    }

    viper.SetConfigFile(filename)
    viper.SetConfigType(CONFIGFILETYPE)

    if err := viper.ReadInConfig(); err != nil {
        log.Printf("read config file error(%v)", err)
        panic(err)
    }

    AppConfig = initAppConfig()
    DBConfig = initDbConfig()

    // validator v9 config register // TODO 这里处理不好，应该将语言配置为全局的，而不是每次用的时候 get 一下
    Uni = ut.New(zh.New(), zh.New())

    viper.SetDefault("LANGUAGE", "zh")
    validateTrans, _ := Uni.GetTranslator(viper.GetString("LANGUAGE"))

    Validate = validator.New()
    _ = zhTranslations.RegisterDefaultTranslations(Validate, validateTrans)
}
