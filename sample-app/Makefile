all: test build pb
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
	@export GO_ENV=development && go build ./cmd/recom-engine && ./recom-engine

dev-perso:
	@@export GO_ENV=development && go build ./cmd/perso-engine && ./perso-engine

test:
	@go test -race ./cmd/recom-engine && \
	@go test -race ./cmd/perso-engine

.PHONY: pb
pb:
	@$(MAKE) -C pkg/grpc/recom-engine
