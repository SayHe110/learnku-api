package main

import (
    "github.com/spf13/pflag"
    "learnku-api/bootstrap"
    "learnku-api/config"
    "learnku-api/http"
)

func main() {
    pflag.Parse()

    config.Init("")

    bootstrap.SetupDB()
    defer bootstrap.DB.CloseDB()

    http.Init(config.C)
}
