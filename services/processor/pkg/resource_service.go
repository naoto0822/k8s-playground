package pkg

import (
	"github.com/naoto0822/k8s-playground/proto/go/processorpb"

	"google.golang.org/grpc"
)

type ResourceService struct {
	cli processorpb.ProcessorSericeClient
}

func NewResourceService() (*ResourceService, error) {
	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	cli := processorpb.NewProcessorSericeClient(conn)
	return &ResourceService{
		cli: cli,
	}, nil
}
