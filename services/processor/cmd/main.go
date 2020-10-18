package main

import (
	"net"

	"github.com/naoto0822/k8s-playground/proto/go/processorpb"
	"github.com/naoto0822/k8s-playground/services/processor/pkg"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger := initLogger()
	logger.Info("processor")

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		logger.Fatal("failed listen", zap.Error(err))
	}

	gs := grpc.NewServer()
	server := pkg.NewServer()
	processorpb.RegisterProcessorServiceServer(gs, &server)

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
