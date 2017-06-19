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
	if [ "$(msg)" != "" ]; then git commit -am "$(msg)"; else git commit -am "Quick Save"; fi
	git push origin $(BRANCH)

orcid: ot.go cmds/orcid/orcid.go
	env CGO_ENABLED=0 go build -o bin/orcid cmds/orcid/orcid.go


clean:
	if [ -d bin ]; then /bin/rm -fR bin; fi
	if [ -d dist ]; then /bin/rm -fR dist; fi

website:
	./mk-website.bash

publish: website
	./publish.bash


dist/linux-amd64:
	mkdir -p dist/bin
	env GOOS=linux GOARCH=amd64 go build -o dist/bin/orcid cmds/orcid/orcid.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-linux-amd64.zip README.md LICENSE INSTALL.md docs/* how-to/* templates/* etc/* bin/*
	rm -fR dist/bin

dist/windows-amd64:
	mkdir -p dist/bin
	env GOOS=windows GOARCH=amd64 go build -o dist/bin/orcid.exe cmds/orcid/orcid.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-windows-amd64.zip README.md LICENSE INSTALL.md docs/* how-to/* templates/* etc/* bin/*
	rm -fR dist/bin

dist/macosx-amd64:
	mkdir -p dist/bin
	env GOOS=darwin GOARCH=amd64 go build -o dist/bin/orcid cmds/orcid/orcid.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-macosx-amd64.zip README.md LICENSE INSTALL.md docs/* how-to/* templates/* etc/* bin/*
	rm -fR dist/bin

dist/raspbian-arm7:
	mkdir -p dist/bin
	env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/orcid cmds/orcid/orcid.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-raspbian-arm7.zip README.md LICENSE INSTALL.md docs/* how-to/* templates/* etc/* bin/*
	rm -fR dist/bin

distribute_docs:
	mkdir -p dist
	cp -v README.md dist/
	cp -v LICENSE dist/
	cp -v INSTALL.md dist/
	if [ -d docs ]; then mkdir -p dist/docs; cp -v docs/*.md dist/docs/; fi
	if [ -d templates ]; then mkdir -p dist/templates; cp -v templates/*.tmpl dist/templates/; fi
	if [ -d etc ]; then mkdir -p dist/etc; cp -v etc/*-example dist/etc/; fi
	if [ -d how-to ]; then mkdir -p dist/how-to; cp -v how-to/*.md dist/how-to/; fi

release: distribute_docs dist/linux-amd64 dist/windows-amd64 dist/macosx-amd64 dist/raspbian-arm7

