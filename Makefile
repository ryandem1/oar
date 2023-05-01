clean:
	docker-compose down --remove-orphans
	rm -Rf dbData

build:
	docker-compose build oar-service

service:
	docker-compose -f docker-compose.yaml -f docker-compose.local.yaml up -d oar-service

test-service:
	cd service; go test -cover

test-pytest-plugin:
	cd pytest; coverage run --omit="*/test*"  -m pytest -s tests; coverage report; rm .coverage

db:
	docker-compose -f docker-compose.yaml -f docker-compose.local.yaml up -d oar-postgres;
	docker-compose -f docker-compose.yaml -f docker-compose.local.yaml up -d wait-for-db;
	docker-compose -f docker-compose.yaml -f docker-compose.local.yaml rm -sfv wait-for-db;
	PGPASSWORD=postgres psql -h localhost -U postgres -d oar -f scripts/sql/init_postgres.sql;


.PHONY: service
