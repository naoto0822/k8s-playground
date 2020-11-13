k8s-dashboard-apply:
	kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.0.0-beta8/aio/deploy/recommended.yaml

get-default-token:
	kubectl -n kube-system get secret | grep default-token | cut -d " " -f 1 | xargs kubectl -n kube-system describe secret

deploy-local:
	kubectl apply -f manifests/resource
	kubectl apply -f manifests/processor
	kubectl apply -f manifests/api

delete-local:
	kubectl delete -f manifests/api
	kubectl delete -f manifests/processor
	kubectl delete -f manifests/resource
