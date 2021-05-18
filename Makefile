PROJECT = plot-maker

GO = GOFLAGS=-mod=vendor go
GOBUILD = $(GO) build $(LDFLAGS)

.PHONY: all
all:	clean build

.PHONY: clean
clean:
	rm -fR build

.PHONY: build
build:
	GOARCH=amd64 $(GOBUILD) -tags static -o build/$(PROJECT) $(PROJECT)/cmd/$(PROJECT)

.PHONY: docker
docker:
	docker build -t plot-maker:latest .
