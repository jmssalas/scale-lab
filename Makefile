.PHONY: tidy up down rate logs experiment-01

tidy:
	go mod tidy

up:
	docker compose up -d

down:
	docker compose down -v

