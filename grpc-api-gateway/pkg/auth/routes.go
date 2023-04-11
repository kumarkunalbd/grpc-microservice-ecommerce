package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/kumarkunalbd/grpc-api-gateway/pkg/auth/routes"
	"github.com/kumarkunalbd/grpc-api-gateway/pkg/config"
)

func RegisterRoutes(gEngine *gin.Engine, conf *config.Config) *ServiceClient {
	serClient := &ServiceClient{
		Client: IntiServiceClient(conf),
	}

	routes := gEngine.Group("/auth")
	routes.POST("/register")
	routes.POST("/login", serClient.Login)

	return serClient

}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}
