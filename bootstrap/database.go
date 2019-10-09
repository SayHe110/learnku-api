package bootstrap

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "learnku-api/config"
    "log"
)

type Database struct {
    Self *gorm.DB
}

var DB *Database

func SetupDB() {
    sourceDB := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
        config.C.DBConfig.Username,
        config.C.DBConfig.Password,
        config.C.DBConfig.Addr,
        config.C.DBConfig.Name,
        true,
        "Local")

    newDB, err := gorm.Open("mysql", sourceDB)
    if err != nil {
        log.Printf("database connection failed. (%v)", err)
        panic(err)
    }

    // 设置数据库连接
    setupDBConfig(newDB)

    DB = &Database{
        Self: newDB,
    }
}

func setupDBConfig(db *gorm.DB) {
    db.LogMode(true)
    db.DB().SetMaxOpenConns(1000) // 最大打开的连接数
    db.DB().SetMaxIdleConns(0)    // 用于设置闲置的连接数
}

func (db *Database) CloseDB() {
    err := db.Self.Close()
    if err != nil {
        log.Printf("database close failed. (%v)", err)
    }
}
