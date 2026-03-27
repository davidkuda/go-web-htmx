# migrations with CLI: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

PG_DSN = postgres://${DB_USER}:${DB_PASSWORD}@${DB_ADDRESS}/${DB_NAME}?sslmode=disable
PG_DSN = ${PG_DSN_ADMIN}

db/migrate/newsql:
	@migrate create \
	-seq \
	-ext=.sql \
	-dir=./migrations \
	${name}

db/migrate/up-roles:
	@if [ -z "$$PG_DSN_ADMIN" ]; then \
	echo "PG_DSN_ADMIN is empty"; \
	exit 1; \
	fi
	@migrate \
	-path=./migrations \
	-database=${PG_DSN_ADMIN} \
	-verbose \
	up

db/migrate/up-all:
	@migrate \
	-path=./migrations \
	-database=${PG_DSN} \
	up

db/migrate/version:
	migrate \
	-path=./migrations/ \
	-database=${PG_DSN} \
	version

# force V: Set version V but don't run migration (ignores dirty state)
db/migrate/force:
	@migrate \
	-path=./migrations/ \
	-database=${PG_DSN} \
	force ${version}

# migrate down one step
db/migrate/down-1:
	@migrate \
	-path=./migrations/ \
	-database=${PG_DSN} \
	down 1

# migrate up one step
db/migrate/up-1:
	@migrate \
	-verbose \
	-path=./migrations/ \
	-database=${PG_DSN} \
	up 1

