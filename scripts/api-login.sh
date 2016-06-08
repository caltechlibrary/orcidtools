#!/bin/bash
#
# This script will display the ORCID_ACCESS_TOKEN based on
# authenticating against the API. You will then need to set that value as the
# environment vairable ORCID_ACCESS_TOKEN to use the other scripts in this directory.
#

function requireSoftware() {
    APP=$(which $1)
    if [ "$APP" = "" ]; then
        echo "Missing $1, $2"
        exit 1
    fi
}

function requireEnvVar() {
    if [ "$2" = "" ]; then
        echo "Missing environment variable: $1"
        exit 1
    fi
}

requireSoftware "curl" "usually installed with your operating system or OS's package manager"
requireSoftware "jq" "See: https://stedolan.github.io/jq/"
requireEnvVar "ORCID_API_URL" $ORCID_API_URL
requireEnvVar "ORCID_CLIENT_ID" $ORCID_CLIENT_ID
requireEnvVar "ORCID_CLIENT_SECRET" $ORCID_CLIENT_SECRET

export ORCID_ACCESS_TOKEN=$(curl -L -H "Accept: application/json" \
    -d "client_id=$ORCID_CLIENT_ID" \
    -d "client_secret=$ORCID_CLIENT_SECRET" \
    -d "scope=/read-public" \
    -d "grant_type=client_credentials" \
    "$ORCID_API_URL/oauth/token" | jq .access_token)
echo 
if [ "$ORCID_ACCESS_TOKEN" != "" ]; then
   echo 
   echo "export ORCID_ACCESS_TOKEN=$ORCID_ACCESS_TOKEN"
   echo 
else
    echo "Login failed $?"
fi
