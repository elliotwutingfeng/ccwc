build_cli:
	go build -o ./dist/ccwc -ldflags "-X 'github.com/elliotwutingfeng/ccwc/cmd/ccwc.version=v0.0.1'" ./cmd/main.go

tests:
	go test -v -race -covermode atomic -coverprofile coverage.out && go tool cover -html coverage.out -o coverage.html
