docker-build:
	docker build -t k8s-playground-api .

docker-push:
	docker login
	docker tag k8s-playground-api:latest naoto0822/k8s-playground-api
	docker push naoto0822/k8s-playground-api
