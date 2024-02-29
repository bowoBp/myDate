include .env

migrate_up:
	migrate -database ${PSQL_MIGRATION_URL} -path pkg/db/migration/sql up


migrate_down:
	migrate -database ${PSQL_MIGRATION_URL} -path pkg/db/migration/sql down 1

fix_migrate_psql_dirty:
	migrate --path pkg/db/migration/sql -database ${PSQL_MIGRATION_URL} force 1

generate_mock:
	mockery --all --keeptree --case underscore --with-expecter