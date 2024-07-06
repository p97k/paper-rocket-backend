DB_CONTAINER_NAME=server-postgres-1
DB_NAME=paper_rocket
POSTGRES_VERSION=15-alpine
POSTGRES_USER=root
POSTGRES_PASSWORD=password
DB_PORT=5433

postgres:
	docker exec -it $(DB_CONTAINER_NAME) psql -U $(POSTGRES_USER) $(DB_NAME)
createdb:
	docker exec -it $(DB_CONTAINER_NAME)  --username=$(POSTGRES_USER) --owner=$(POSTGRES_USER)
dropdb:
	docker exec -it $(DB_CONTAINER_NAME) dropdb
up:
	docker-compose up
upd:
	docker-compose up -d
buildup:
	docker-compose up --build

.PHONY: postgres createdb dropdb run buildup upd