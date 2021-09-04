.PHONY: run deps migration-create migration-up migration-down

DB_DSN=postgresql://postgres:password@db/twitch2slack?sslmode=disable

migration-create: # accepts parameter name={migration-name}
	docker-compose run migrations create -ext sql -dir migrations $(name)

migration-up: # accepts parameter count={number-of-migrations}
	docker-compose run migrations -path=/migrations -database=$(DB_DSN) up $(count)

migration-down: # accepts parameter count={number-of-migrations}
	docker-compose run migrations -path=/migrations -database=$(DB_DSN) down $(count)

deps:
	docker-compose up -d db migrations

run:
	go run cmd/t2s/*.go
