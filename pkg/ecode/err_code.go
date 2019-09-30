package ecode

var (
    InvalidTokenParams = NewECode(1000, "缺少token参数")
    ParseTokenError    = NewECode(1001, "解析token出错")
    TokenExpireTime    = NewECode(1002, "token以过期")
    RefreshTokenError  = NewECode(1003, "token以过期")
)
