reset-dummy-data:
	rm alive.db && sqlite3 < migration.sql && sqlite3 < ./example/dummy_data.sql

test:
	go test ./...