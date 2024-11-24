create_migration:
	migrate create -ext sql -dir ./migrations -seq init

.PHONY: create_migration

migrate_up:
	migrate -path ./migrations -database 'postgres://david:davidpswrd@localhost:5436/putlist_db?sslmode=disable' up

.PHONY: migrate_up

migrate_down:
	migrate -path ./migrations -database 'postgres://david:davidpswrd@localhost:5436/putlist_db?sslmode=disable' down

.PHONY: migrate_down

create_migration_test:
	migrate create -ext sql -dir ./tests/migrations -seq init_test

.PHONY: create_migrarion_test

migrate_up_test:
	migrate -path ./tests/migrations -database 'postgres://david:davidpswrd@localhost:5436/putlist_db_test?sslmode=disable' up

.PHONY: migrate_up_test

migrate_down_test:
	migrate -path ./tests/migrations -database 'postgres://david:davidpswrd@localhost:5436/putlist_db_test?sslmode=disable' down

.PHONY: migrate_down_test