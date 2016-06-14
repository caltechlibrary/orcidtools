#!/bin/bash
#
# This script will search an ORCID profile by providing a ORCID number
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
    export ORCID_NUMBER="$1"
fi

requireEnvVar "ORCID_API_URL" $ORCID_API_URL
requireEnvVar "ORCID_ACCESS_TOKEN" $ORCID_ACCESS_TOKEN
requireEnvVar "ORCID_NUMBER" $ORCID_NUMBER

curl -L -H "Content-Type: $OUT_FORMAT" \
    -H "Authorization: Bearer $ORCID_ACCESS_TOKEN" \
    -X GET "$ORCID_API_URL/v1.2/$ORCID_NUMBER/orcid-profile"
