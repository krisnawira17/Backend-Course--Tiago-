build:
	@go build -o bin/go-backend-learn cmd/main.go

test:
	@go test -v ./...
run: build
	@./bin/go-backend-learn