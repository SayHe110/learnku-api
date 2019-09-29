package config

import (
    "github.com/spf13/viper"
    "time"
)

type appConfig struct {
    Name       string
    Runmode    string
    Addr       string
    Port       int
    Url        string
    JWTSecret  string
    ExpireTime time.Duration
}

func initAppConfig() *appConfig {
    viper.SetDefault("APP.NAME", "default name")
    viper.SetDefault("APP.RUNMODE", "debug")
    viper.SetDefault("APP.ADDR", "127.0.0.1")
    viper.SetDefault("APP.PORT", "8080")
    viper.SetDefault("APP.URL", "127.0.0.1:8080")
    viper.SetDefault("JWT.SECRET", "abc")
    viper.SetDefault("JWT.EXPIRETIME", 1)

    return &appConfig{
        Name:       viper.GetString("APP.NAME"),
        Runmode:    viper.GetString("APP.RUNMODE"),
        Addr:       viper.GetString("APP.ADDR"),
        Port:       viper.GetInt("APP.PORT"),
        Url:        viper.GetString("APP.URL"),
        JWTSecret:  viper.GetString("JWT.SECRET"),
        ExpireTime: viper.GetDuration("JWT.EXPIRETIME") * time.Hour,
    }
}
