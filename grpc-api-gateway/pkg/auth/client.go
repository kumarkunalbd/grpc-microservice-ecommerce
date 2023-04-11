package auth

import (
	"fmt"

	"github.com/kumarkunalbd/grpc-api-gateway/pkg/auth/pb"
	"github.com/kumarkunalbd/grpc-api-gateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func IntiServiceClient(conf *config.Config) pb.AuthServiceClient {

	clinetC, err := grpc.Dial(conf.AuthSvcUrl, insecure.NewCredentials())

	if err != nil {
		fmt.Println("Counld not connect:", err)
	}

	return pb.NewAuthServiceClient(clinetC)

}
