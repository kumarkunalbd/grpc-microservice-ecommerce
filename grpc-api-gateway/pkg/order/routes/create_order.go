package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kumarkunalbd/grpc-api-gateway/pkg/order/pb"
)

type CreateOrderRequestBody struct {
	ProductId int64
	Quantity  int64
}

func CreateOrder(ctx *gin.Context, c pb.OrderServiceClient) {
	requestBody := CreateOrderRequestBody{}

	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId, _ := ctx.Get("userId")
	res, err := c.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		ProductId: requestBody.ProductId,
		Quantity:  requestBody.Quantity,
		UserId:    userId.(int64),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
