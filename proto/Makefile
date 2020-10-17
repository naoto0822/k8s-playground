protoc:
	protoc --go_out=plugins=grpc:. ./src/*
	rm -rf ./go
	mv ./github.com/naoto0822/k8s-playground/proto/go ./go
	rm -rf ./github.com
