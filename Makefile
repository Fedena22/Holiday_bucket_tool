test: fumpt ##Run tests
	go install github.com/mfridman/tparse@latest
	go test ./... -cover -json | tparse -all
build: fumpt generate ##Build the app
	@mkdir -p out/bin
	CGO_ENABLED=0 go build -ldflags="-w -s" -o out/bin ./...

lint: fumpt ##Run linit
	go install github.com/kisielk/errcheck@latest
	${GOPATH}/bin/errcheck ./...
	go install honnef.co/go/tools/cmd/staticcheck@latest
	${GOPATH}/bin/staticcheck -checks all ./...
fumpt: ##Run gofumpt
	go install mvdan.cc/gofumpt@latest
	${GOPATH}/bin/gofumpt -l -w .
generate: fumpt
	go generate ./...

run:
	go run ./cmd/server.go