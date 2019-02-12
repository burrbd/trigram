build: dep
	go build -race ./cmd/web/web.go

test:
	go test -race -cover ./...

dep:
	go dep ensure
