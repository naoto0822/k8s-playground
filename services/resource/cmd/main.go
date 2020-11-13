package main

import (
	"net"

	"github.com/naoto0822/k8s-playground/proto/go/resourcepb"
	"github.com/naoto0822/k8s-playground/services/resource/pkg"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger := initLogger()
	logger.Info("start resource service")

	lis, err := net.Listen("tcp", ":10003")
	if err != nil {
		logger.Fatal("failed listen", zap.Error(err))
	}

	gs := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_zap.UnaryServerInterceptor(logger),
		),
	)
	server := pkg.NewServer()
	resourcepb.RegisterResourceServiceServer(gs, server)

	logger.Info("start resource server")
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
