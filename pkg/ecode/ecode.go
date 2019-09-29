package ecode

type ecode struct {
    Code    int
    Message string
}

var (
    INVALID_TOKEN_PARAMS = &ecode{5000, "缺少token参数"}
    PARSE_TOKEN_ERROR    = &ecode{5001, "解析token出错"}
    TOKEN_EXPIRE_TIME    = &ecode{5002, "token以过期"}
)
