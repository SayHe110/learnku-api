package ecode

var (
    InvalidTokenParams = NewECode(5000, "缺少token参数")
    ParseTokenError    = NewECode(5001, "解析token出错")
    TokenExpireTime    = NewECode(5002, "token以过期")
)
