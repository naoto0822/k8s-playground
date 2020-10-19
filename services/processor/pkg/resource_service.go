package pkg

import (
	"context"

	"github.com/naoto0822/k8s-playground/proto/go/resourcepb"
	"github.com/naoto0822/k8s-playground/proto/go/types"

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

func (s *ResourceService) GetRamenList(ctx context.Context) ([]*types.Ramen, error) {
	req := &resourcepb.GetRamenRequest{}
	res, err := s.cli.GetRamenList(ctx, req)
	if err != nil {
		return nil, err
	}

	return res.Ramens, nil
}
