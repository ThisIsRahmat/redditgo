# //This is to indicate to MAke that our commands don't create any files
.PHONY: postgress adminer migrate

postgres:
	docker run --rm -ti  -p 5432:5432 --network redditgo -e POSTGRES_PASSWORD=secret postgres

adminer:
	docker run --rm -ti -p 5430:8080  --network redditgo  adminer

migrate: 
	docker run --rm -v $(PWD)/migrations:/migrations --network redditgo migrate/migrate -path=/migrations -database "postgres://postgres:secret@172.23.0.2:5432/postgres?sslmode=disable" up
