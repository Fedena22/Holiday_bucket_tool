test: fumpt ##Run tests
	go install github.com/mfridman/tparse@latest
	go test ./... -cover -json | tparse -all
build: fumpt generate ##Build the app
	@mkdir -p out/bin
	CGO_ENABLED=0 go build -ldflags="-w -s" -o out/bin ./...

lint: fumpt ##Run linit
	go install github.com/kisielk/errcheck@latest
	errcheck ./...
	go install honnef.co/go/tools/cmd/staticcheck@latest
	staticcheck -checks all ./...
fumpt: ##Run gofumpt
	go install mvdan.cc/gofumpt@latest
	gofumpt -l -w .
generate: fumpt
	go generate ./...

run: generate
	go run ./cmd/server.go