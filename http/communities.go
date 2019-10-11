package http

import (
    "github.com/gin-gonic/gin"
    "learnku-api/model/communities"
    "learnku-api/pkg/response"
    "net/http"
)

func communityList(ctx *gin.Context) {
    res, err := communitiesSvc.GetList();
    if err != nil {
        response.JSON(ctx, http.StatusOK, "获取列表失败", err)
        return
    }

    response.JSON(ctx, http.StatusOK, "获取列表成功", res)
}

func communityUpdate(ctx *gin.Context) {
    comm := &communities.CommunityParams{}
    if err := ctx.ShouldBind(&comm); err != nil {
        response.JSON(ctx, 50001, "参数绑定失败", err)
        return
    }

    if err := comm.UpdateCommunityValidator().Errors; err != nil {
        response.JSON(ctx, 50002, "参数验证失败", comm.UpdateCommunityValidator())
        return
    }

    if err := communitiesSvc.Update(comm); err != nil {
        response.JSON(ctx, http.StatusOK, "更新失败", err)
        return
    }

    response.JSON(ctx, http.StatusOK, "更新成功", nil)
}
