clean:
	docker-compose down --remove-orphans
	rm -Rf dbData

build:
	docker-compose build oar-service

service:
	docker-compose -f docker-compose.yaml -f docker-compose.local.yaml up -d oar-service

test:
	cd service; go test -cover

db:
	docker-compose -f docker-compose.yaml -f docker-compose.local.yaml up -d oar-postgres


.PHONY: service
