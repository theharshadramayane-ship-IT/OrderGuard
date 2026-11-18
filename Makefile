.PHONY: run test build migrate-up migrate-down

run:
	cd backend && go run main.go

test:
	cd backend && go test -v ./...

build:
	cd backend && go build -o orderguard.exe main.go
	cd rust/orderguard-safety && cargo build --release

migrate-up:
	migrate -path backend/migrations -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path backend/migrations -database "$(DATABASE_URL)" down 1