.PHONY: dbup dbdown reset apiup dbshell test

dbup:
	docker compose up -d db

dbdown:
	docker compose down db

reset:
	docker compose down -v

apiup:
	docker compose up --build api

dbshell:
	docker compose exec db psql -U postgres

test:
	docker compose up --build test
