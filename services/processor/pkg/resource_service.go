package pkg

import (
	"github.com/naoto0822/k8s-playground/proto/go/resourcepb"

	"google.golang.org/grpc"
)

type ResourceService struct {
	cli resourcepb.ResourceServiceClient
}

func NewResourceService() (*ResourceService, error) {
	conn, err := grpc.Dial("127.0.0.1:8082", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	cli := resourcepb.NewResourceServiceClient(conn)
	return &ResourceService{
		cli: cli,
	}, nil
}
