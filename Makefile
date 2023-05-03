# Removes all database data and tears down all apps
clean:
	docker-compose down --remove-orphans
	rm -Rf dbData

# Builds the oar-service
build-service:
	docker-compose build oar-service

# Builds the enrich-ui image
build-enrich-ui:
	docker-compose build oar-enrich-ui

# Runs the oar-service container
service:
	docker-compose -f docker-compose.yaml -f docker-compose.local.yaml up -d oar-service

# Runs the enrich-ui container
enrich-ui:
	docker-compose -f docker-compose.yaml -f docker-compose.local.yaml up -d oar-enrich-ui

# Includes hot-reloading and runs on localhost instead of Docker
enrich-ui-dev:
	cd enrich-ui; npm dev run

# Runs the unit tests on the oar-service
test-service:
	cd service; go test -cover

# Runs the unit tests on the pytest-oar plugin
test-pytest-plugin:
	cd pytest; coverage run --omit="*/test*"  -m pytest -s -vv tests; coverage report; rm .coverage

# Runs the playwright e2e tests on the enrich UI. Must run enrich-ui first to run e2e tests. Tests are performed
# on "productionalized" enrich-ui docker container
test-enrich-ui-e2e:
	cd enrich-ui; npm run test

test-enrich-ui-unit:
	cd enrich-ui; npm run test:unit

# Starts the oar-postgres database, waits for the database boot-up and runs the db init script
db:
	docker-compose -f docker-compose.yaml -f docker-compose.local.yaml up -d oar-postgres;
	docker-compose -f docker-compose.yaml -f docker-compose.local.yaml up -d wait-for-db;
	docker-compose -f docker-compose.yaml -f docker-compose.local.yaml rm -sfv wait-for-db;
	PGPASSWORD=postgres psql -h localhost -U postgres -d oar -f scripts/sql/init_postgres.sql;


.PHONY: service enrich-ui
