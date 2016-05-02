#
# Simple Makefile for conviently testing, building and deploying experiment.
#
build: ot.go
	go build
	go build -o bin/orcidapi cmds/orcidapi/orcidapi.go
	go build -o bin/orcidpdr2db cmds/orcidpdr2db/orcidpdr2db.go

test:
	go test

clean:
	if [ -d bin ]; then rm bin/*; fi
