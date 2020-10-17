package pkg

import (
	"context"

	"github.com/naoto0822/k8s-playground/proto/go/processorpb"
)

type Server struct {
}

func NewServer() Server {
	return Server{}
}

func (s Server) GetSobaList(context.Context, *processorpb.GetSobaRequest) (*processorpb.GetSobaResponse, error) {
	sobas := []*processorpb.Soba{
		{
			Id:   1,
			Name: "arimasa",
		},
	}

	res := &processorpb.GetSobaResponse{
		Sobas: sobas,
	}

	return res, nil
}
