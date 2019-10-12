package cache

import (
    "github.com/gomodule/redigo/redis"
    "learnku-api/config"
    "time"
)

func RedisCache() *redis.Pool {
    pwd := config.C.RediesConfig.Password

    return &redis.Pool{
        MaxIdle:     config.C.RediesConfig.Maxidle,
        MaxActive:   config.C.RediesConfig.Maxactive,
        IdleTimeout: config.C.RediesConfig.Idletimeout,
        Dial: func() (conn redis.Conn, e error) {
            c, err := redis.Dial("tcp", config.C.RediesConfig.Addr)
            if err != nil {
                return nil, err
            }

            // 密码
            if pwd != "" {
                if _, err := c.Do("AUTH", pwd); err != nil {
                    _ = c.Close()
                    return nil, err
                }
            }

            return c, nil
        },
        TestOnBorrow: func(c redis.Conn, t time.Time) error {
            _, err := c.Do("PING")

            return err
        },
    }
}
