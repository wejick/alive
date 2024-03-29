server-version:=`cat ./cmd/alive-server/version.txt`
agent-version:=`cat ./cmd/alive-agent/version.txt`

reset-dummy-data:
	-rm alive.db
	sqlite3 < migration.sql && sqlite3 < ./example/dummy_data.sql

test:
	go test --count=1 ./...

run-server:
	go build ./cmd/alive-server
	./alive-server --dbpath ./alive.db

run-agent:
	go build ./cmd/alive-agent
	./alive-agent --configpath ./config.json

run-linter:
	podman run --rm -v $(pwd):/app -v ~/.cache/golangci-lint/v1.50.1:/root/.cache -w /app golangci/golangci-lint:v1.50.1 golangci-lint run -v --timeout 5m

generate-mock:
	mockery --all --keeptree
	-rm -r mocks/internal_mock
	mv mocks/internal mocks/internal_mock

build-container-server:
	podman build -t wejick/alive-server:$(server-version) -f dockerfile.server .

build-container-agent:
	podman build -t wejick/alive-agent:$(agent-version) -f dockerfile.agent .

