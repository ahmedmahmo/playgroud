postgres:
	docker-compose up -d
postgres_down:
	docker-compose down
create_db:
	docker exec -it postgres createdb --username postgres bank
drop_db:
	docker exec -it postgres dropdb --username postgres bank
migrate_up:
	migrate -path migration -database "postgresql://postgres:password@localhost:5432/bank?sslmode=disable" -verbose up
migrate_down:
	migrate -path migration -database "postgresql://postgres:password@localhost:5432/bank?sslmode=disable" -verbose down


.PHONY: create_db drop_db postgres postgres_down migrate_up migrate_down