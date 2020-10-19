package pkg

import (
	"context"

	"github.com/naoto0822/k8s-playground/proto/go/processorpb"
)

type Server struct {
	resourceService *ResourceService
}

func NewServer(resourceService *ResourceService) *Server {
	return &Server{
		resourceService: resourceService,
	}
}

func (s *Server) GetRamenList(ctx context.Context, _ *processorpb.GetRamenRequest) (*processorpb.GetRamenResponse, error) {
	ramens, err := s.resourceService.GetRamenList(ctx)
	if err != nil {
		return nil, err
	}

	res := &processorpb.GetRamenResponse{
		Ramens: ramens,
	}
	return res, nil
}
