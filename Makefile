.PHONY: tidy up down run-api

ifneq (,$(wildcard ./.env))
    include .env
    export
endif

tidy:
	go mod tidy

up:
	docker compose -f deploy/compose.yml up --build -d

down:
	docker compose -f deploy/compose.yml down -v

run-api:
	go run cmd/api/main.go
