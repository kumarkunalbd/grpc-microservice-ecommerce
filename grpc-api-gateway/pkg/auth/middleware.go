package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kumarkunalbd/grpc-api-gateway/pkg/auth/pb"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svClient *ServiceClient) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{svc: svClient}
}

func (c *AuthMiddlewareConfig) AuthRequired(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("authorization")

	if strings.TrimSpace(authorization) == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	authToken := strings.Split(authorization, "Bearer ")
	if len(authToken) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res, err := c.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: authToken[1],
	})

	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	ctx.Set("userId", res.UserId)

	ctx.Next()

}
