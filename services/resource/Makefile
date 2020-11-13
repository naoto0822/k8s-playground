dep:
	go get ./...

docker-build:
	docker build -t k8s-playground-resource .

docker-push: docker-build
	docker login
	docker tag k8s-playground-resource:latest naoto0822/k8s-playground-resource
	docker push naoto0822/k8s-playground-resource
