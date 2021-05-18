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
	GOARCH=amd64 GOOS=linux $(GOBUILD) -tags static -o build/$(PROJECT) $(PROJECT)/cmd/$(PROJECT)
