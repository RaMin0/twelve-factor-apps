run:
	env $$(cat config.env | xargs) go run main.go
.PHONY: run

run_db:
	docker run --rm -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=app -d postgres:alpine
.PHONY: run_db

populate_db:
	cat dump.sql | docker exec -i $$(docker ps -q) psql -h localhost -U root app
.PHONY: populate_db
