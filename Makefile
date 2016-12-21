#
# ORCID tools Makefile
#
PROJECT = ot

VERSION = $(shell grep 'Version = ' $(PROJECT).go | cut -d \" -f 2)

BRANCH = $(shell git branch | grep '* ' | cut -d\  -f 2)

PROG_LIST = orcid

build: $(PROG_LIST)

test:
	go test

lint:
	golint ot.go
	golint cmds/orcid/orcid.go

install:
	env GOBIN=$(HOME)/bin go install cmds/orcid/orcid.go


save:
	git commit -am "Quick Save"
	git push origin $(BRANCH)

orcid: ot.go cmds/orcid/orcid.go
	env CGO_ENABLED=0 go build -o bin/orcid cmds/orcid/orcid.go


clean:
	if [ -d bin ]; then /bin/rm -fR bin; fi
	if [ -d dist ]; then /bin/rm -fR dist; fi
	if [ -f $(PROJECT)-$(VERSION)-release.zip ]; then /bin/rm $(PROJECT)-$(VERSION)-release.zip; fi

release:
	./mk-release.bash

website:
	./mk-website.bash

publish: website
	./publish.bash
