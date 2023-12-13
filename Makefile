CONTAINER_NAME := wegonice-db
CONTAINER_PORT := 5432
DB_NAME := wegonice
DB_USER := root
DB_USER_PASSWORD := secret

# Docker commands #
connect-to-database:
	docker exec -it $(CONTAINER_NAME) psql -U $(DB_USER) $(DB_NAME)

create-container:
	docker run --name $(CONTAINER_NAME) -p $(CONTAINER_PORT):$(CONTAINER_PORT) -e POSTGRES_USER=$(DB_USER) -e POSTGRES_PASSWORD=$(DB_USER_PASSWORD) -d postgres:16-alpine

destroy-container:
	docker stop $(CONTAINER_NAME) && docker rm $(CONTAINER_NAME)	

start-container:
	docker start $(CONTAINER_NAME)

tail-container-logs:
	docker logs -f $(CONTAINER_NAME)

.PHONY: connect-to-database create-container destroy-container start-container tail-container-logs

# Database management #
create-db:
	docker exec -it $(CONTAINER_NAME) createdb --username=$(DB_USER) --owner=$(DB_USER) $(DB_NAME)

dropdb:
	docker exec -it $(CONTAINER_NAME) dropdb $(DB_NAME)

migrate-down:
	 migrate --path db/migration --database "postgresql://$(DB_USER):$(DB_USER_PASSWORD)@localhost:$(CONTAINER_PORT)/$(DB_NAME)?sslmode=disable" --verbose down

migrate-down1:
	 migrate --path db/migration --database "postgresql://$(DB_USER):$(DB_USER_PASSWORD)@localhost:$(CONTAINER_PORT)/$(DB_NAME)?sslmode=disable" --verbose down 1

migrate-up:
	 migrate --path db/migration --database "postgresql://$(DB_USER):$(DB_USER_PASSWORD)@localhost:$(CONTAINER_PORT)/$(DB_NAME)?sslmode=disable" --verbose up

migrate-up1:
	 migrate --path db/migration --database "postgresql://$(DB_USER):$(DB_USER_PASSWORD)@localhost:$(CONTAINER_PORT)/$(DB_NAME)?sslmode=disable" --verbose up 1

.PHONY: create-db dropdb migrate-down migrate-down1 migrate-up1

# SQLC #
sqlc:
	sqlc generate

.PHONY: sqlc

# Testing #
test:
	go test -v --cover ./...

.PHONY: test

# Server #
server:
	go run main.go

.PHONY: server