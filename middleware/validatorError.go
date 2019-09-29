package middleware

import (
    "github.com/gin-gonic/gin"
    "github.com/go-playground/universal-translator"
    "gopkg.in/go-playground/validator.v9"
    "net/http"
)

type ErrorHandler struct {
    uni *ut.UniversalTranslator
}

func NewValidatorErrorHandler(uni *ut.UniversalTranslator) *ErrorHandler {
    return &ErrorHandler{
        uni: uni,
    }
}

func (h *ErrorHandler) HandlerValidatorError(c *gin.Context) {
    c.Next()

    errorToPrint := c.Errors.ByType(gin.ErrorTypePublic).Last()
    if errorToPrint != nil {
        if errs, ok := errorToPrint.Err.(validator.ValidationErrors); ok {
            trans, _ := h.uni.GetTranslator("zh") // 这里也可以通过获取 HTTP Header 中的 Accept-Language 来获取用户的语言设置
            c.JSON(http.StatusBadRequest, gin.H{
                "errors": errs.Translate(trans),
            })
            return
        }
    }
}
