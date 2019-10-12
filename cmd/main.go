package main

import (
    "github.com/spf13/pflag"
    "learnku-api/config"
    "learnku-api/http"
    "learnku-api/pkg/casbin"
)

func main() {
    pflag.Parse()

    config.Init("")

    casbin.Casbin()

    http.Init(config.C)
}
