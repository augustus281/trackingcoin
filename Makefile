run:
	go run cmd/main.go 
migrateup:
	migrate -path database/migration -database "postgresql://root:mysecurepassword@localhost:5433/trackingcoin?sslmode=disable" -verbose up
migratedown:
	migrate -path database/migration -database "postgresql://root:mysecurepassword@localhost:5433/trackingcoin?sslmode=disable" -verbose down
sqlc:
	sqlc generate
.PHONY: run migrateup migratedown