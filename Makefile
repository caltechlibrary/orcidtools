#
# Simple Makefile for conviently testing, building and deploying experiment.
#
build: ot.go cmds/ot/ot.go cmds/otpdr/otpdr.go
	go build -o bin/ot cmds/ot/ot.go
	go build -o bin/otpdr cmds/otpdr/otpdr.go

test:
	go test

clean:
	if [ -d bin ]; then rm bin/*; fi
