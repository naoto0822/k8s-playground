package main

import (
	"net"

	"github.com/naoto0822/k8s-playground/proto/go/processorpb"
	"github.com/naoto0822/k8s-playground/services/processor/pkg"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger := initLogger()
	logger.Info("processor")

	lis, err := net.Listen("tcp", ":10002")
	if err != nil {
		logger.Fatal("failed listen", zap.Error(err))
	}

	server, err := initServer()
	if err != nil {
		logger.Fatal("failed init server", zap.Error(err))
	}

	gs := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_zap.UnaryServerInterceptor(logger),
		),
	)
	processorpb.RegisterProcessorServiceServer(gs, server)

	logger.Info("start processor server")
	if err := gs.Serve(lis); err != nil {
		logger.Fatal("failed grpc server", zap.Error(err))
	}
}

func initLogger() *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	return logger
}

func initServer() (*pkg.Server, error) {
	resourceService, err := pkg.NewResourceService()
	if err != nil {
		return nil, err
	}

	server := pkg.NewServer(resourceService)
	return server, nil
}
