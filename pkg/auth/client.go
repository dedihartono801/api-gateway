package auth

import (
	"fmt"
	"log"

	"github.com/dedihartono801/api-gateway/pkg/config"
	pb "github.com/dedihartono801/protobuf/auth/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	opts := []grpc.DialOption{}
	tls := true

	if tls {
		certFile := "../ssl/auth-svc/ca.crt"

		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Error while loading CA trust certificates: %v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	// using WithSecure() because no SSL running
	cc, err := grpc.Dial(c.AuthSvcUrl, opts...)

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAuthServiceClient(cc)
}
