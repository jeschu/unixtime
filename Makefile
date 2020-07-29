GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)

default: travis

clean:
	@rm -rf build

install-gox:
	go get github.com/mitchellh/gox

travis: clean install-gox
	gox -output="build/{{.OS}}/{{.Arch}}/unixtime"

build:
	gox -osarch="$(GOOS)/$(GOARCH)" -output="build/{{.OS}}/{{.Arch}}/unixtime"