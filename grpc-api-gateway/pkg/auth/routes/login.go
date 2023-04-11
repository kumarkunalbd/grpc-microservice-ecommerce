package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kumarkunalbd/grpc-api-gateway/pkg/auth/pb"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context, c pb.AuthServiceClient) {

	loginBody := LoginRequestBody{}

	if err := ctx.BindJSON(&loginBody); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Email:    loginBody.Email,
		Password: loginBody.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
