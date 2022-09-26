reset-dummy-data:
	rm alive.db && sqlite3 < migration.sql && sqlite3 < ./example/dummy_data.sql

test:
	go test ./...

run-server:
	go build ./cmd/alive-server
	./alive-server

run-agent:
	go build ./cmd/alive-agent
	./alive-agent