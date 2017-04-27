PROJECT=github.com/fulldump/logo
GOPATH=$(shell pwd)
GO=go
GOCMD=GOPATH=$(GOPATH) GO15VENDOREXPERIMENT=1 $(GO)


.PHONY: all
all: dependencies test

.PHONY: clean
clean:
	rm -fr src/
	rm -fr bin/
	rm -fr pkg/

.PHONY: setup
setup:
	mkdir -p src/$(PROJECT)
	rm -fr src/$(PROJECT)
	ln -s ../../.. src/$(PROJECT)

.PHONY: dependencies
dependencies: setup
	$(GOCMD) get -t $(PROJECT)

.PHONY: test
test:
	$(GOCMD) test $(PROJECT)

.PHONY: coverage
coverage:
	$(GOCMD) test ./src/$(PROJECT) -cover -covermode=count -coverprofile=coverage.out; \
	$(GOCMD) tool cover -html=coverage.out
