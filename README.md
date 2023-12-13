# wegonice-api

## Initialize database

- Start database container with `make create-container`
- Create database inside container with `make create-db`
- Migrate schema with `make migrate-up`

## Database handling

### Create schema

- I used `dbdiagram.io` for this project
- [schema](https://dbdiagram.io/d/wegonice-db-6579938456d8064ca0f24284)

### Add more migrations

```zsh
migrate create -ext sql -dir ./db/migration -seq <migration-name>
```
