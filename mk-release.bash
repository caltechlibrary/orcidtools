#!/bin/bash
#
# Make releases for Linux/amd64, Linux/ARM6 and Linux/ARM7 (Raspberry Pi), Windows, and Mac OX X (darwin)
#

VERSION=$(grep 'Version = ' ot.go | cut -d\" -f 2)
RELEASE_NAME=ot-$VERSION
echo "Preparing $RELEASE_NAME-release.zip"
for PROGNAME in orcid ; do
  echo "Cross compiling $PROGNAME"
  env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/$PROGNAME cmds/$PROGNAME/$PROGNAME.go
  env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/$PROGNAME cmds/$PROGNAME/$PROGNAME.go
  env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -o dist/raspberrypi-arm6/$PROGNAME cmds/$PROGNAME/$PROGNAME.go
  env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspberrypi-arm7/$PROGNAME cmds/$PROGNAME/$PROGNAME.go
  env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/$PROGNAME.exe cmds/$PROGNAME/$PROGNAME.go
done

# Copy docs and templates
for ITEM in README.md INSTALL.md LICENSE scripts templates htdocs/index.md htdocs/css htdocs/js; do
    if [ -f $ITEM ] || [ -d $ITEM ]; then
        echo "cp -vR $ITEM dist/"
        cp -vR $ITEM dist/
    else
        echo "Skipping $ITEM, not found"
    fi
done

# Copy configuration examples
for ITEM in etc/ot.bash-example; do
    if [ ! -d dist/etc ]; then
        mkdir -p dist/etc
    fi
    cp -vR $ITEM dist/etc/
done

echo "Zipping $RELEASE_NAME-release.zip"
zip -r "$RELEASE_NAME-release.zip" dist/*
