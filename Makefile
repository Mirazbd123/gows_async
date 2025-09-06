db_login:
	psql ${DB_URL}

db_create_migration:
	migrate create -ext sql -dir db/migrations -seq ${name}