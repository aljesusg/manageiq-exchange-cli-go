# Project specific variables
PROJECT=manageiq-exchange
# --- the rest of the file should not need to be configured ---

# GO env
GOPATH=$(shell pwd)
GO=go
GOCMD=GOPATH=$(GOPATH) $(GO)
RELEASE_PATH := ${GOPATH}/release
DISTPATH := bin/$(PROJECT)

# Build versioning
COMMIT = $(shell git log -1 --format="%h" 2>/dev/null || echo "0")
VERSION=$(shell git describe --tags --always)
BUILD_DATE = $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
FLAGS = -ldflags "\
  -X constants.COMMIT=$(COMMIT) \
  -X constants.VERSION=$(VERSION) \
  -X constants.BUILD_DATE=$(BUILD_DATE) \
  "

GOBUILD = $(GOCMD) build $(FLAGS)

.PHONY: all
all:	build

.PHONY: build
build: format test compile

.PHONY: compile
compile:
	GOARCH=amd64 GOOS=darwin $(GOBUILD) -o $(DISTPATH).darwin ./src/$(PROJECT)
	GOARCH=amd64 GOOS=linux $(GOBUILD) -o $(DISTPATH).linux ./src/$(PROJECT)

.PHONY: deploy
deploy: coverage build
	echo "Creating tar file"
	tar -zcf $(RELEASE_PATH)/$(PROJECT)-$(VERSION).tar.gz bin/$(PROJECT)*

.PHONY: format
format:
	@for gofile in $$(find ./$(PROJECT) -name "*.go"); do \
		echo "formatting" $$gofile; \
		gofmt -w $$gofile; \
	done

.PHONY: run
run:
	- $(GOCMD) run ./main.go

.PHONY: test
test:
	$(GOCMD) test -v -race -tags safe ./...

.PHONY: coverage
coverage:
		rm -fr coverage
		mkdir -p coverage
		$(GOCMD) list $(PROJECT)/... > coverage/packages
		@i=a ; \
		while read -r P; do \
			i=a$$i ; \
			$(GOCMD) test ./src/$$P -cover -coverpkg $$P -covermode=count -coverprofile=coverage/$$i.out; \
		done <coverage/packages
		echo "mode: count" > coverage/coverage
		cat coverage/*.out | grep -v "mode: count" >> coverage/coverage
		$(GOCMD) tool cover -html=coverage/coverage

.PHONY: CI-Coverage
CI-Coverage: coverage
		go get -v github.com/mattn/goveralls
		goveralls -coverprofile=coverage/coverage -service=travis-ci -covermode=count
.PHONY: clean
clean:
	rm -fR bin pkg
