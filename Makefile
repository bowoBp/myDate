include .env

migrate_up:
	migrate -database ${PSQL_MIGRATION_URL} -path pkg/db/migration/sq up


migrate_down:
	migrate -database ${PSQL_MIGRATION_URL} -path pkg/db/migration/sq down 1