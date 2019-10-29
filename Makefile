default: travis

travis:
	@rm -rf build
	go get github.com/mitchellh/gox
	gox -output="build/{{.OS}}/{{.Arch}}/unixtime"