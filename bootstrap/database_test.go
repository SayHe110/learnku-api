package bootstrap

import (
    "learnku-api/config"
    "testing"
)

func init() {
    config.Init("../.env.yaml")
}

func TestSetupDB(t *testing.T) {
    //
}
