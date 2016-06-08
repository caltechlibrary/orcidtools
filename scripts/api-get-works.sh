#!/bin/bash
#
# This script will a list of ORCID works by providing a ORCID number
# email address.
#

function requireEnvVar() {
    if [ "$2" = "" ]; then
        echo "Missing $1"
        exit 1
    fi
}


if [ "$1" != "" ]; then
    export ORCID_NUMBER="$1"
fi
requireEnvVar "ORCID_API_URL" $ORCID_API_URL
requireEnvVar "ORCID_ACCESS_TOKEN" $ORCID_ACCESS_TOKEN
requireEnvVar "ORCID_NUMBER" $ORCID_NUMBER

curl -L -H "Content-Type: application/vdn.orcid+xml" \
    -H "Authorization: Bearer $ORCID_ACCESS_TOKEN" \
    -X GET "$ORCID_API_URL/v1.2/$ORCID_NUMBER/orcid-works"
