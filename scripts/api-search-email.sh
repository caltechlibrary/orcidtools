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

if [ "$1" != "" ]; then
    export EMAIL="$1"
fi
requireEnvVar "ORCID_API_URL" $ORCID_API_URL
requireEnvVar "ORCID_ACCESS_TOKEN" $ORCID_ACCESS_TOKEN
requireEnvVar "EMAIL" $EMAIL

curl -L -H "Content-Type: application/vdn.orcid+xml" \
    -H "Authorization: Bearer $ORCID_ACCESS_TOKEN" \
    -X GET "$ORCID_API_URL/v1.2/search/orcid-bio/?q=email:$EMAIL"


