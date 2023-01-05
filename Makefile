clean:
	docker-compose down

build:
	docker-compose build oar-service

run:
	docker-compose up -d oar-service
