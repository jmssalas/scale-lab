.PHONY: tidy up down

tidy:
	go mod tidy

up:
	docker compose -f deploy/compose.yml up --build -d

down:
	docker compose -f deploy/compose.yml down -v
