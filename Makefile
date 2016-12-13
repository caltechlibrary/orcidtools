#
# ORCID tools Makefile
#
PROJECT = ot
PROG_LIST = orcid

build: $(PROG_LIST)

test:
	go test

lint:
	golint ot.go
	golint cmds/orcid/orcid.go

install:
	go install cmds/orcid/orcid.go

save:
	git commit -am "Quick Save"
	git push origin master

orcid: ot.go cmds/orcid/orcid.go
	env CGO_ENABLED=0 go build -o bin/orcid cmds/orcid/orcid.go


clean:
	if [ -d bin ]; then /bin/rm -fR bin; fi
	if [ -d dist ]; then /bin/rm -fR dist; fi
	if [ -f $(PROJECT)-release.zip ]; then /bin/rm $(PROJECT)-release.zip; fi

release:
	./mk-release.bash

website:
	./mk-website.bash

publish: website
	./publish.bash
