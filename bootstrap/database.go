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
        config.DBConfig.Username,
        config.DBConfig.Password,
        config.DBConfig.Addr,
        config.DBConfig.Name,
        true,
        "Local")

    DB, err := gorm.Open("mysql", sourceDB)
    if err != nil {
        log.Printf("database connection failed. (%v)", err)
        panic(err)
    }

    // 设置数据库连接
    setupDB(DB)
}

func setupDB(db *gorm.DB) {
    db.LogMode(true)
    db.DB().SetMaxOpenConns(1000) // 最大打开的连接数
    db.DB().SetMaxIdleConns(0)    // 用于设置闲置的连接数
}

func (db *Database) CloseDB() {
    err := db.Self.Close()
    if err !=nil {
        log.Printf("database close failed. (%v)", err)
    }
}
