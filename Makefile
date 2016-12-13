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
	env CGO_ENABLED=0 GOBIN=$(HOME)/bin go install cmds/orcid/orcid.go

save:
	git commit -am "Quick Save"
	git push origin master

orcid: ot.go cmds/orcid/orcid.go
	env CGO_ENABLED=0 go build -o bin/orcid cmds/orcid/orcid.go


