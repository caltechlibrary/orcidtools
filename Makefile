#
# Simple Makefile for conviently testing, building and deploying experiment.
#
build: ot.go cmds/ot/ot.go cmds/otpdr/otpdr.go
	go build -o bin/ot cmds/ot/ot.go
	go build -o bin/otpdr cmds/otpdr/otpdr.go

install:
	env GOBIN=$(HOME)/bin go install cmds/ot/otpdr.go
	env GOBIN=$(HOME)/bin go install cmds/ot/otpdr.go

test:
	go test

clean:
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi
	if [ -d ot-binary-release.zip ]; then rm -f ot-binary-release.zip; fi

release:
	./mk-release.sh

