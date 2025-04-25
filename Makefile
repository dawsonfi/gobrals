configure-proxy:
	go env -w GOPROXY=direct

test:
	go test ./...

test_v:
	go test -v ./...

test_short:
	go test ./... -short

test_race:
	go test ./... -short -race

test_stress:
	go test -tags=stress -timeout=30m ./...

test_codecov:
	go test -coverprofile=coverage.out -covermode=atomic ./...

build:
	go build ./...

vet:
	go vet ./...

fmt:
	go fmt ./...
	goimports -l -w .

lint:
	golangci-lint run

run:
	go run .