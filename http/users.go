package http

import (
    "github.com/gin-gonic/gin"
    "learnku-api/pkg/response"
    userSvc "learnku-api/service/users"
    "net/http"
)

func userList(c *gin.Context) {
    res, _ := userSvc.Users()
    response.JSON(c, http.StatusOK, "", res)
}
