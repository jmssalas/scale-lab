.PHONY: tidy up down run-api

tidy:
	go mod tidy

up:
	docker compose up -d

down:
	docker compose down -v

run-api:
	go run cmd/api/main.go
