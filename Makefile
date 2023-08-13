help:: ## Show this help
	@fgrep -h "##" $(MAKEFILE_LIST) | sort | fgrep -v fgrep | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

storage-up: ## Initialize db in Kubernetes
	kubectl apply -f db-persistent-volume.yaml
	kubectl apply -f db-volume-claim.yaml
	kubectl apply -f db-configmap.yaml
	kubectl apply -f db-deployment.yaml
	kubectl apply -f db-service.yaml

api-up: ## run api
	kubectl apply -f api.yaml

app-up: ## run app
	kubectl apply -f app.yaml