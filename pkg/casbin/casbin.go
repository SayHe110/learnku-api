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

var (
    ResourcesRuleName = &resourcesRuleName{}
    CasbinEnforce     *casbin.Enforcer
)

type resourcesRuleName struct {
    Communities string
    Users       string
    Topics      string
    Categories  string
}

func init() {
    ResourcesRuleName = &resourcesRuleName{
        Communities: "communities",
        Users:       "users",
        Topics:      "topics",
        Categories:  "categories",
    }
}

func InitCasbinEnforce() {
    CasbinEnforce, _ = casbinEnforceFun()
}

func Enforce(sub, obj, act string) (ok bool, err error) {
    err = CasbinEnforce.LoadPolicy()
    if err != nil {
        return false, errors.New("casbin 加载策略失败~")
    }

    ok, err = CasbinEnforce.Enforce(sub, obj, act)
    return
}

func casbinEnforceFun() (en *casbin.Enforcer, err error) {
    adapterGorm, err := gormadapter.NewAdapterByDB(db.NewMysql(config.C))
    if err != nil {
        return nil, errors.New("新建 Gorm 适配器失败~")
    }

    en, err = casbin.NewEnforcer(_casbinAuthConfigFile, adapterGorm)
    if err != nil {
        return nil, errors.New("casbin 执行失败~")
    }

    return
}
