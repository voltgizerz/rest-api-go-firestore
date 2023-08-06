run:
	go run ./cmd/main.go

format:
	gofmt -w ./..
	revive -config revive.toml ./...

build:
	go build ./cmd/main.go

test:
	go test -v -cover ./...

cover-out:
	go test ./...  -coverpkg=./... -coverprofile ./coverage.out