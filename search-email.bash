#!/bin/bash
#
# This script will display the ORCID_ACCESS_TOKEN based on
# authenticating against the API. You will then need to set that value as the
# environment vairable ORCID_ACCESS_TOKEN to use the other scripts in this directory.
#

function requireSoftware() {
	APP="$(which "$1")"
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
requireSoftware "jsoncols" "See: https://caltechlibrary.github.io/datatools/"
requireEnvVar "ORCID_API_URL" "${ORCID_API_URL}"
requireEnvVar "ORCID_CLIENT_ID" "${ORCID_CLIENT_ID}"
requireEnvVar "ORCID_CLIENT_SECRET" "${ORCID_CLIENT_SECRET}"

if [ "$1" = "" ]; then
    cat<<EOF

USAGE: $(basename "$0") EMAIL_ADDRESS

Example: 

Search for all the people with email ending in 'example.edu'

    $(basrname "$0") '*@*example.edu'"

EOF
	exit 1
fi

ACCESS_TOKEN=$(curl --silent -L -H "Accept: application/json" \
	-d "client_id=${ORCID_CLIENT_ID}" \
	-d "client_secret=${ORCID_CLIENT_SECRET}" \
	-d "scope=/read-public" \
	-d "grant_type=client_credentials" \
	"${ORCID_API_URL}/oauth/token")

if [ "$?" = "0" ]; then
	ORCID_ACCESS_TOKEN=$(echo "${ACCESS_TOKEN}" | jsoncols .access_token)
	export ORCID_ACCESS_TOKEN
else
	echo "Login failed $?"
	echo "--> $ACCESS_TOKEN"
	exit 1
fi

#
# This script will search for an ORCID by providing a partial or complete
# email address.
#
#echo "Access token: ${ORCID_ACCESS_TOKEN}"
OUT_FORMAT="application/json"
EMAIL="$1"

curl --silent -L -H "Content-Type: ${OUT_FORMAT}" \
	-H "Authorization: Bearer ${ORCID_ACCESS_TOKEN}" \
	-X GET "${ORCID_API_URL}/v2.0/search/?q=email:${EMAIL}" |\
    jsoncols --quiet '.result[:]["orcid-identifier"].path' |\
    jsonrange -values
