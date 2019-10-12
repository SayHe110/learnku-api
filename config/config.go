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

type Config struct {
    AppConfig    *appConfig
    DBConfig     *dbConfig
    RediesConfig *redisConfig
    Uni          *ut.UniversalTranslator
    Validate     *validator.Validate
}

var C = &Config{}

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

    appConfig := initAppConfig()
    dbConfig := initDbConfig()
    rediesConfig := initRedisConfig()

    // validator v9 config register // TODO 这里处理不好，应该将语言配置为全局的，而不是每次用的时候 get 一下
    uni := ut.New(zh.New(), zh.New())

    viper.SetDefault("LANGUAGE", "zh")
    validateTrans, _ := uni.GetTranslator(viper.GetString("LANGUAGE"))

    validate := validator.New()
    _ = zhTranslations.RegisterDefaultTranslations(validate, validateTrans)

    C = &Config{
        AppConfig:    appConfig,
        DBConfig:     dbConfig,
        RediesConfig: rediesConfig,
        Uni:          uni,
        Validate:     validate,
    }
}
