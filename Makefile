include .env

migrate_up:
	migrate -database ${PSQL_MIGRATION_URL} -path pkg/db/migration/sql up


migrate_down:
	migrate -database ${PSQL_MIGRATION_URL} -path pkg/db/migration/sql down 1