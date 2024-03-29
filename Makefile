db-up:
	docker-compose  up --remove-orphans --build -d
run.vendor:
	go mod vendor
run.swag:
	~/go/bin/swag init -g ./cmd/documentation/main.go
run.swagger:
	go run cmd/documentation/main.go
run.product:
	go run cmd/api/service.product/main.go
run.transaction:
	go run cmd/api/service.transaction/main.go

