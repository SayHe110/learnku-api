package config

import (
    "github.com/spf13/viper"
    "log"
)

const (
    CONFIGFILENAME = "./.env.yaml"
    CONFIGFILETYPE = "yaml"
)

var (
    AppConfig *appConfig
    DBConfig  *dbConfig
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
}
