.PHONY: tidy up down run-api

ifneq (,$(wildcard ./.env))
    include .env
    export
endif

tidy:
	go mod tidy

up:
	docker compose up -d

down:
	docker compose down -v

run-api:
	go run cmd/api/main.go
