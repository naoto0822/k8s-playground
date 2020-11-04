package pkg

import (
	"context"

	"github.com/naoto0822/k8s-playground/proto/go/processorpb"
	"github.com/naoto0822/k8s-playground/proto/go/types"

	"google.golang.org/grpc"
)

type ProcessorService struct {
	cli processorpb.ProcessorServiceClient
}

func NewProcessorService() (*ProcessorService, error) {
	conn, err := grpc.Dial("127.0.0.1:10002", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	cli := processorpb.NewProcessorServiceClient(conn)
	return &ProcessorService{
		cli: cli,
	}, nil
}

func (s *ProcessorService) GetSobaList(ctx context.Context) ([]*types.Ramen, error) {
	req := &processorpb.GetRamenRequest{}
	res, err := s.cli.GetRamenList(ctx, req)
	if err != nil {
		return nil, err
	}

	return res.Ramens, nil
}
