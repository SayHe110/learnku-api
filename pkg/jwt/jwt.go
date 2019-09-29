package jwt

import "learnku-api/config"

var jwtSecret = []byte(config.AppConfig.JWTSecret)


