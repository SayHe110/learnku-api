package main

import (
    "github.com/spf13/pflag"
    "learnku-api/config"
    "learnku-api/http"
)

func main() {
    pflag.Parse()

    config.Init("")

    http.Init()
}
