all: test build
clean:
	@rm ./recom-engine
	@rm ./perso-engine

build:
	@go build -race ./cmd/recom-engine && \
	@go build -race ./cmd/perso-engine

build-recom:
	@go build ./cmd/recom-engine

build-perso:
	@go build ./cmd/perso-engine

dev-recom:
	@go build ./cmd/recom-engine && ./recom-engine

dev-perso:
	@go build ./cmd/perso-engine && ./perso-engine

test:
	@go test -race ./cmd/recom-engine && \
	@go test -race ./cmd/perso-engine