BINARY=service
build:
	go build ${LDFLAGS} -o ${BINARY} cmd/web/*.go
test:
	go test -json > report.json -cover -coverprofile=coverage.out -race ./...
web: build
	./${BINARY} -E dev
