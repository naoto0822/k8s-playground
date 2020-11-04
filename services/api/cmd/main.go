package main

import (
	"github.com/naoto0822/k8s-playground/services/api/pkg"

	"github.com/labstack/echo"
	"go.uber.org/zap"
)

func main() {
	logger := initLogger()
	logger.Info("start api service")

	_, err := pkg.InitConfig()
	if err != nil {
		logger.Fatal("failed init config", zap.Error(err))
	}

	handler, err := initHandler()
	if err != nil {
		logger.Fatal("failed init handler", zap.Error(err))
	}

	e := echo.New()
	v1 := e.Group("/v1")
	v1.GET("/ruok", handler.Ruok)
	v1.GET("/soba", handler.GetSobaList)

	//go func() {
	if err := e.Start(":10001"); err != nil {
		logger.Fatal("failed server start", zap.Error(err))
	}
	//}()
}

func initLogger() *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	return logger
}

func initHandler() (*pkg.Handler, error) {
	processorService, err := pkg.NewProcessorService()
	if err != nil {
		return nil, err
	}

	handler := pkg.NewHandler(processorService)
	return handler, nil
}
