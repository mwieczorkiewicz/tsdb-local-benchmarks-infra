influxup:
	docker compose -f ./infra/influxdb/influxdb-docker-compose.yaml up -d 

influxdown:
	docker compose -f ./infra/influxdb/influxdb-docker-compose.yaml down

postgresup:
	docker compose -f ./infra/postgres/postgres-docker-compose.yaml up -d 

postgresdown:
	docker compose -f ./infra/postgres/postgres-docker-compose.yaml down

timescaleup:
	docker compose -f ./infra/timescale/timescale-docker-compose.yaml up -d 

timescaledown:
	docker compose -f ./infra/timescale/timescale-docker-compose.yaml down

allup: influxup postgresup timescaleup

alldown: influxdown postgresdown timescaledown