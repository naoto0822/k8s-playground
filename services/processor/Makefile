docker-build:
	docker build -t k8s-playground-processor .

docker-push:
	docker login
	docker tag k8s-playground-processor:latest naoto0822/k8s-playground-processor
	docker push naoto0822/k8s-playground-processor
