db-up:
	docker-compose  up --remove-orphans --build -d
run.vendor:
	go mod vendor
run.swagger:
	go run cmd/doc/main.go
run:
	go run cmd/api/main.go
swag:
	~/go/bin/swag init -g ./cmd/doc/main.go
