.PHONY: all

service := golangservicel0

start-servers:
	@echo "  >  Building binary..."
	docker-compose up -d postgresql

init-postgresql: start-servers
	docker-compose exec -T postgresql psql -U default < db/dbimage/pgsql.sql

connect-to-postgresql: start-servers
	docker-compose exec postgresql psql -U default

build:
	docker-compose build --no -cache ${service}

start: start-servers
	docker-compose --compatibility up --build -d ${service}

run:
	docker-compose --compatibility up --build ${service}

clean:
	docker-compose down --volumes --remove-orphans
