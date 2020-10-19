package pkg

import (
	"context"

	"github.com/naoto0822/k8s-playground/proto/go/resourcepb"
	"github.com/naoto0822/k8s-playground/proto/go/types"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) GetRamenList(_ context.Context, _ *resourcepb.GetRamenRequest) (*resourcepb.GetRamenResponse, error) {
	ramens := []*types.Ramen{
		{
			Id:   1,
			Name: "arimasa",
		},
	}

	res := &resourcepb.GetRamenResponse{
		Ramens: ramens,
	}

	return res, nil
}
