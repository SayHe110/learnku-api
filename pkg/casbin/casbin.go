package casbin

import (
    "github.com/casbin/casbin/v2"
    "github.com/casbin/gorm-adapter/v2"
    _ "github.com/go-sql-driver/mysql"
    "learnku-api/config"
    "learnku-api/pkg/db"
    "log"
)

func Casbin() {
    a, _ := gormadapter.NewAdapterByDB(db.NewMysql(config.C))
    e, err := casbin.NewEnforcer("./pkg/casbin/auth_model.conf", a)

    if err != nil {
        log.Panic(err)
    }

    if err := e.LoadPolicy(); err != nil {
        log.Panic(err)
    }

    _, _ = e.Enforce("alice", "data1", "read")

    _ = e.SavePolicy()
}
