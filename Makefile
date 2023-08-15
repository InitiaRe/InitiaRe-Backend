run:
	go run main.go

test:
	go test -cover ./...

migrate:
	go run migration/migrate.go

tidy:
	go mod tidy

doc:
	echo "Starting swagger generating"
	swag fmt
	swag init -g main.go --pd

