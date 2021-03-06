#!/bin/bash
#
# This script will search for an ORCID by providing a partial or complete 
# email address.
#

function requireEnvVar() {
    if [ "$2" = "" ]; then
        echo "Missing $1"
        exit 1
    fi
}

#OUT_FORMAT="application/vdn.orcid+xml"
OUT_FORMAT="application/json"

if [ "$1" != "" ]; then
    export QTERM="$1"
fi

requireEnvVar "ORCID_API_URL" $ORCID_API_URL
requireEnvVar "ORCID_ACCESS_TOKEN" $ORCID_ACCESS_TOKEN
requireEnvVar "QTERM" $QTERM

curl -L -H "Content-Type: $OUT_FORMAT" \
    -H "Authorization: Bearer $ORCID_ACCESS_TOKEN" \
    -X GET "$ORCID_API_URL/v2.0/search/?q=$QTERM"




