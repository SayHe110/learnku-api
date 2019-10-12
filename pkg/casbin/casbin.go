package casbin

import (
    "errors"
    "github.com/casbin/casbin/v2"
    "github.com/casbin/gorm-adapter/v2"
    _ "github.com/go-sql-driver/mysql"
    "learnku-api/config"
    "learnku-api/pkg/db"
)

const (
    _casbinAuthConfigFile = "./pkg/casbin/auth_model.conf"
)

func Enforce(sub, obj, act string) (ok bool, err error) {
    adapterGorm, err := gormadapter.NewAdapterByDB(db.NewMysql(config.C))
    if err != nil {
        return false, errors.New("新建 Gorm 适配器失败~")
    }

    enforce, err := casbin.NewEnforcer(_casbinAuthConfigFile, adapterGorm)
    if err != nil {
        return false, errors.New("casbin 执行失败~")
    }

    err = enforce.LoadPolicy()
    if err != nil {
        return false, errors.New("casbin 加载策略失败~")
    }

    ok, err = enforce.Enforce(sub, obj, act)
    return
}
