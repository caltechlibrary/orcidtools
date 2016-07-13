#!/bin/bash
#

function checkSoftware() {
    APP=$(which bibfilter)
    if [ "$APP" = "" ]; then
        echo "Missing bibfilter, see https://caltechlibrary.github.io/bibtex for installation info"
        exit 1
    fi
}


checkSoftware
if [ $# -eq 0 ]; then
    bibfilter -exclude=comments
else
    bibfilter -exclude=comments $1 $2 $3 $4 $5 $6 $7 $8 $9
fi

