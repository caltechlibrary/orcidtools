#
# Simple Makefile for conviently testing, building and deploying experiment.
#
build: ot.go
	go build
	go build -o bin/orcidapi cmds/orcidapi/orcidapi.go

test:
	go test

clean:
	if [ -d bin ]; then rm bin/*; fi
