.PHONY: dbup dbdown reset apiup dbshell

dbup:
	docker compose up -d db

dbdown:
	docker compose down db

reset:
	docker compose down -v

apiup:
	docker compsoe up api

dbshell:
	docker compose exec db psql -U postgres
