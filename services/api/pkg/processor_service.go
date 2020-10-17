package pkg

import (
	"context"

	"github.com/naoto0822/k8s-playground/proto/go/processorpb"

	"google.golang.org/grpc"
)

type ProcessorService struct {
	cli processorpb.ProcessorSericeClient
}

func NewProcessorService() (*ProcessorService, error) {
	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	cli := processorpb.NewProcessorSericeClient(conn)
	return &ProcessorService{
		cli: cli,
	}, nil
}

func (s ProcessorService) GetSobaList(ctx context.Context) ([]*processorpb.Soba, error) {
	req := &processorpb.GetSobaRequest{}
	res, err := s.cli.GetSobaList(ctx, req)
	if err != nil {
		return nil, err
	}

	return res.Sobas, nil
}
