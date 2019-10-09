package db

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "learnku-api/config"
    "log"
)

func NewMysql(c *config.Config) (db *gorm.DB) {
    sourceDB := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
        c.DBConfig.Username,
        c.DBConfig.Password,
        c.DBConfig.Addr,
        c.DBConfig.Name,
        true,
        "Local")

    db, err := gorm.Open("mysql", sourceDB)
    if err != nil {
        log.Printf("database connection failed. (%v)", err)
        panic(err)
    }

    // 设置数据库连接
    setupDBConfig(db)

    return
}

func setupDBConfig(db *gorm.DB) {
    db.LogMode(true)
    db.DB().SetMaxOpenConns(1000) // 最大打开的连接数
    db.DB().SetMaxIdleConns(0)    // 用于设置闲置的连接数
}
