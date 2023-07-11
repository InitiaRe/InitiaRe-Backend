run:
	go run main.go

test:
	go test -cover ./...

migrate:
	go run migrations/migrate.go

tidy:
	go mod tidy

doc:
	echo "Starting swagger generating"
	swag fmt
	swag init -g cmd/main.go --pd

