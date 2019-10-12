package config

import (
    "github.com/spf13/viper"
    "time"
)

type dbConfig struct {
    Addr     string
    Name     string
    Username string
    Password string
}

type redisConfig struct {
    Addr        string
    Password    string
    Maxidle     int
    Maxactive   int
    Idletimeout time.Duration
}

func initDbConfig() *dbConfig {
    viper.SetDefault("DB.ADDR", "127.0.0.1:3306")
    viper.SetDefault("DB.NAME", "DatabaseName")
    viper.SetDefault("DB.USERNAME", "name")
    viper.SetDefault("DB.PASSWORD", "password")

    return &dbConfig{
        Addr:     viper.GetString("DB.ADDR"),
        Name:     viper.GetString("DB.NAME"),
        Username: viper.GetString("DB.USERNAME"),
        Password: viper.GetString("DB.PASSWORD"),
    }
}

func initRedisConfig() *redisConfig {
    viper.SetDefault("REDIS.ADDR", "127.0.0.1:6379")
    viper.SetDefault("REDIS.PASSWORD", "")
    viper.SetDefault("REDIS.MAXIDLE", 10)
    viper.SetDefault("REDIS.MAXACTIVE", 0)
    viper.SetDefault("REDIS.IDLETIMEOUT", 0)

    return &redisConfig{
        Addr:        viper.GetString("REDIS.ADDR"),
        Password:    viper.GetString("REDIS.PASSWORD"),
        Maxidle:     viper.GetInt("REDIS.MAXIDLE"),
        Maxactive:   viper.GetInt("REDIS.MAXACTIVE"),
        Idletimeout: viper.GetDuration("REDIS.IDLETIMEOUT"),
    }
}
