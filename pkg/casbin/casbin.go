package casbin

import (
    "github.com/billcobbler/casbin-redis-watcher"
    "github.com/casbin/casbin/v2"
    "github.com/casbin/gorm-adapter/v2"
    _ "github.com/go-sql-driver/mysql"
    "learnku-api/config"
    "learnku-api/pkg/db"
    "log"
)

const (
    _casbinAuthConfigFile = "./pkg/casbin/auth_model.conf"
)

func Casbin() {
    w, err := rediswatcher.NewWatcher(config.C.RediesConfig.Addr, rediswatcher.Password(config.C.RediesConfig.Password))

    if err != nil {
        log.Panic(err.Error())
    }

    a, _ := gormadapter.NewAdapterByDB(db.NewMysql(config.C))
    e, err := casbin.NewEnforcer(_casbinAuthConfigFile, a)

    _ = e.SetWatcher(w)
    if err != nil {
        log.Panic(err)
    }

    if err := e.LoadPolicy(); err != nil {
       log.Panic(err)
    }

    _ = w.SetUpdateCallback(func(s string) {
        log.Println(s)
    })

    _, _ = e.Enforce("alice", "data1", "read")

    _ = e.SavePolicy()
}
