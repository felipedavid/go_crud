dsn = "postgres://web:secret@localhost/go_crud?sslmode=disable"

migrateup:
	migrate -path migration -database $(dsn) -verbose up

migratedown:
	migrate -path migration -database $(dsn) -verbose down

.PHONY: migrateup migratedown
