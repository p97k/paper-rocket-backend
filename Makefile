DB_CONTAINER_NAME=postgres15
POSTGRES_VERSION=15-alpine
POSTGRES_USER=root
POSTGRES_PASSWORD=password
DB_PORT=5432

up: down
	docker run --name $(DB_CONTAINER_NAME) -e POSTGRES_USER=$(POSTGRES_USER) -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) -p $(DB_PORT):5432 -d postgres:$(POSTGRES_VERSION)

down:
	@if [ $$(docker ps -a -q -f name=$(DB_CONTAINER_NAME)) ]; then \
		docker stop $(DB_CONTAINER_NAME); \
	fi

postgres:
	docker exec -it $(DB_CONTAINER_NAME) psql
createdb:
	docker exec -it $(DB_CONTAINER_NAME)  --username=$(POSTGRES_USER) --owner=$(POSTGRES_USER)
dropdb:
	docker exec -it $(DB_CONTAINER_NAME) dropdb
migrateup:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5432/paper_rocket?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5432/paper_rocket?sslmode=disable" -verbose down

.PHONY: up postgres createdb dropdb migrateup migratedown down