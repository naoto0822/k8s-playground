package pkg

import (
	"context"

	"github.com/naoto0822/k8s-playground/proto/go/processorpb"
	"github.com/naoto0822/k8s-playground/proto/go/types"
)

type Server struct {
}

func NewServer() Server {
	return Server{}
}

func (s Server) GetRamenList(ctx context.Context, req *processorpb.GetRamenRequest) (*processorpb.GetRamenResponse, error) {
	ramens := []*types.Ramen{
		{
			Id:   1,
			Name: "arimasa",
		},
	}

	res := &processorpb.GetRamenResponse{
		Ramens: ramens,
	}

	return res, nil
}
