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


status:
	git status

save:
	git commit -am "Quick Save"
	git push origin $(BRANCH)

orcid: ot.go cmds/orcid/orcid.go
	env CGO_ENABLED=0 go build -o bin/orcid cmds/orcid/orcid.go


clean:
	if [ -d bin ]; then /bin/rm -fR bin; fi
	if [ -d dist ]; then /bin/rm -fR dist; fi
	if [ -f $(PROJECT)-$(VERSION)-release.zip ]; then /bin/rm $(PROJECT)-$(VERSION)-release.zip; fi

website:
	./mk-website.bash

publish: website
	./publish.bash

release: dist/linux-amd64 dist/windows-amd64 dist/macosx-amd64 dist/raspbian-arm7
	mkdir -p dist
	mkdir -p dist/etc/
	cp -v README.md dist/
	cp -v LICENSE dist/
	cp -v INSTALL.md dist/
	cp -vR scripts dist/
	cp -vR templates dist/
	cp -v etc/*-example dist/etc/
	zip -r $(PROJECT)-$(VERSION)-release.zip dist/*

dist/linux-amd64:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/orcid cmds/orcid/orcid.go

dist/windows-amd64:
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/orcid.exe cmds/orcid/orcid.go

dist/macosx-amd64:
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/orcid cmds/orcid/orcid.go

dist/raspbian-arm7:
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspberrypi-arm7/orcid cmds/orcid/orcid.go


