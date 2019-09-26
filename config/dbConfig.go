package config

import "github.com/spf13/viper"

type dbConfig struct {
    Addr     string
    Name     string
    Username string
    Password string
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
