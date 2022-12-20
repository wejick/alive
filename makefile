server-version:=`cat ./cmd/alive-server/version.txt`

reset-dummy-data:
	rm alive.db && sqlite3 < migration.sql && sqlite3 < ./example/dummy_data.sql

test:
	go test --count=1 ./...

run-server:
	go build ./cmd/alive-server
	./alive-server

run-agent:
	go build ./cmd/alive-agent
	./alive-agent

build-container-server:
	podman build -t wejick/alive-server:$(server-version) -f dockerfile.server .