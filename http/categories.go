package http

import (
    "github.com/gin-gonic/gin"
    categoriesSvc "learnku-api/handler/categories"
    "learnku-api/model/categories"
    "learnku-api/pkg/response"
    "net/http"
)

func categoriesList(ctx *gin.Context) {

}

func categoriesStore(ctx *gin.Context) {
    categoryParams := categories.CategoryParams{}

    if err := ctx.ShouldBind(&categoryParams); err != nil {
        response.JSON(ctx, 50001, "参数绑定错误", nil)
        return
    }

    if err := categoryParams.CategoryStoreValidator().Errors; err != nil {
        response.JSON(ctx, 50002, "参数验证错误", categoryParams.CategoryStoreValidator())
        return
    }

    if err := categoriesSvc.Store(categoryParams); err != nil {
        response.JSON(ctx, 50003, "创建失败", err.Error())
        return
    }

    response.JSON(ctx, http.StatusNoContent, "创建成功", nil)
}
