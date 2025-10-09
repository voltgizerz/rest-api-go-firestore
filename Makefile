MOCK_DIR=internal/app/mocks
INTERFACE_DIR=internal/app/interfaces

run:
	go run ./cmd/app.go

format:
	gofmt -w ./..
	revive -config revive.toml ./...

build:
	go build ./cmd/app.go

test:
	go test -v -cover ./...

cover-out:
	go test ./...  -coverpkg=./... -coverprofile ./coverage.out

install:
	go install go.uber.org/mock/mockgen@latest

gen-mock: install
	@mkdir -p $(MOCK_DIR)
	@for file in $(shell find $(INTERFACE_DIR) -name '*.go'); do \
		filename=$$(basename $$file); \
		mock_file=$(MOCK_DIR)/mocks_$$filename; \
		echo "Generating mock for $$file -> $$mock_file"; \
		mockgen -source=$$file -destination=$$mock_file -package=mocks; \
	done