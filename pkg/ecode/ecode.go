package ecode

type ECode struct {
    Code    int
    Message string
}

var (
    _ecode = map[int]ECode{} // register ecode
)

func NewECode(code int, msg string) ECode {
    eCode := ECode{
        Code:    code,
        Message: msg,
    }

    _ecode[code] = eCode

    return eCode
}

func GetECode(code int) string {
    msg, ok := _ecode[code]
    if !ok {
        return ""
    }

    return msg.Message
}
