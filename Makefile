run:
	go run cmd/main.go & envoy -c api-gateway/api-gateway.yaml
migrateup:
	migrate -path database/migration -database "postgresql://root:mysecurepassword@localhost:5432/trackingcoin?sslmode=disable" -verbose up
migratedown:
	migrate -path database/migration -database "postgresql://root:mysecurepassword@localhost:5432/trackingcoin?sslmode=disable" -verbose down
migrateforce:
	migrate -path database/migration -database "postgresql://root:mysecurepassword@localhost:5432/trackingcoin?sslmode=disable" force 1
sqlc:
	sqlc generate --file=database/sqlc.yaml
compose:
	docker compose up -d
wire:
	cd internal/wire && wire
envoy:
	envoy -c api-gateway/api-gateway.yaml
deploy:
	kubectl apply -f deployments/deployment.yaml
deploy_service:
	kubectl apply -f deployments/service.yaml
delete:
	kubectl delete -f deployments/deployment.yaml
delete_service:
	kubectl delete -f deployments/service.yaml
build_docker:
	docker build -t myapp -f build/Dockerfile .
run_docker:
	docker run -d -p 80:80 myapp
.PHONY: run migrateup migratedown sqlc compose wire envoy