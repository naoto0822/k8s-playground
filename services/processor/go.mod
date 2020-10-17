module github.com/naoto0822/k8s-playground/services/processor

go 1.15

replace github.com/naoto0822/k8s-playground/proto => ../../proto

require (
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/ktr0731/evans v0.9.1 // indirect
	github.com/naoto0822/k8s-playground/proto v0.0.0-00010101000000-000000000000
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0
	golang.org/x/net v0.0.0-20201016165138-7b1cca2348c0 // indirect
	golang.org/x/sys v0.0.0-20201017003518-b09fb700fbb7 // indirect
	google.golang.org/genproto v0.0.0-20201015140912-32ed001d685c // indirect
	google.golang.org/grpc v1.33.0
	google.golang.org/protobuf v1.25.0 // indirect
)
