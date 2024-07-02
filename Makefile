DB_CONTAINER_NAME=postgres15
POSTGRES_VERSION=15-alpine
POSTGRES_USER=root
POSTGRES_PASSWORD=password
DB_PORT=5432

postgres:
	docker exec -it $(DB_CONTAINER_NAME) psql
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