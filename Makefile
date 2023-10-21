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

gen-mock:
	mockgen -source=./internal/app/interfaces/user.interface.go -destination=internal/app/mocks/mocks_user.interface.go
	mockgen -source=./internal/app/interfaces/auth.interface.go -destination=internal/app/mocks/mocks_auth.interface.go