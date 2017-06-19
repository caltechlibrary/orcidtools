
# ORCID Exploration and Integration

## Get All Caltech ORCID ids

+ [Getting ORCIDs by institution](http://members.orcid.org/finding-orcid-record-holders-your-institution) - get orcids with @caltech.edu in the name
+ [Tutorial on Search for ORCIDs](http://members.orcid.org/api/tutorial-searching-api-12-and-earlier) -
+ [Tutorial on searching the API with Curl](http://members.orcid.org/api/tutorial-retrieve-data-public-api-curl-12-and-earlier)
+ [How do I get public data](http://support.orcid.org/knowledgebase/articles/223698)

### Steps

1. Get an Access token (does not expire)

```shell
    curl -i -L -H "Accept: application/json" -d "client_id=APP-01XX65MXBF79VJGF" -d "client_secret=3a87028d-c84c-4d5f-8ad5-38a93181c9e1" -d "scope=/read-public" -d "grant_type=client_credentials" "https://pub.sandbox.orcid.org/oauth/token"
```

```json

    {"access_token":"39a1d9c5-e753-41b8-b676-a82142ef67ae","token_type":"bearer","refresh_token":"a3a2420a-b964-4cee-bc15-7bdd873b1643","expires_in":631138518,"scope":"/read-public","orcid":null}
```

2. Using access tokens search by email domain

```shell
curl -H "Content-Type: application/orcid+xml" -H "Authorization: 39a1d9c5-e753-41b8-b676-a82142ef67ae" "https://pub.sandbox.orcid.org/v2.0/search/record/?q=email:*@caltech.edu
```
3. Save data
4. transform into XLSX/CSV files

### Algorithm

1. For each faculty name in faculty list matching eprints
2. For each DOI associated with an EPrint query the Orcid Public API for associated Name and DOI, if hit remember the information and ORCID
3. For each faculty member without a related DOI query the Orcid Public API by name, if only one record return remember the orcid and return any associated data
4. For each EPrint with ORCID make sure we've have added the name to faculty list, to get full details query the Orcid Public API and save record

## Activities around API

### Recipe for authenticating with API

This Bash script uses *curl* and [jq](https://stedolan.github.io/jq/) to interact with the ORCID API.  
It requires the following environment variables to be set.

+ ORCID_API_URL
+ ORCID_CLIENT_ID
+ ORCID_SECRET

If successful it will display a string you can copy and paste into your shell session that exports ORCID_ACCESS_TOKEN. This last
environment variable is used by the other script examples.

```shell
    #!/bin/bash
    #
    # This script will display the ORCID_ACCESS_TOKEN based on
    # authenticating against the API. You will then need to set that value as the
    # environment vairable ORCID_ACCESS_TOKEN to use the other scripts in this directory.
    #
    
    function requireEnvVar() {
        if [ "$2" = "" ]; then
            echo "Missing environment variable: $1"
            exit 1
        fi
    }
    
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
```

### Recipe for searching for ORCID by email

This script will search for an email address (or a wildcarded email address). It requires the following
environment variable to be set.

+ ORCID_API_URL
+ ORCID_ACCESS_TOKEN
+ EMAIL (can all be passed as an option to the script)

```shell
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
        -X GET "$ORCID_API_URL/v2.0/search/record/?q=email:$EMAIL"
```

### Recipe for getting an ORCID profile

This script will get an ORCID profile. It requires the following environment variable to be set.

+ ORCID_API_URL
+ ORCID_ACCESS_TOKEN
+ ORICD_NUMBER (can be provided as an option to the script)

```shell
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
    
    
    if [ "$1" != "" ]; then
        export ORCID_NUMBER="$1"
    fi
    requireEnvVar "ORCID_API_URL" $ORCID_API_URL
    requireEnvVar "ORCID_ACCESS_TOKEN" $ORCID_ACCESS_TOKEN
    requireEnvVar "ORCID_NUMBER" $ORCID_NUMBER
    
    curl -L -H "Content-Type: application/vdn.orcid+xml" \
        -H "Authorization: Bearer $ORCID_ACCESS_TOKEN" \
        -X GET "$ORCID_API_URL/v2.0/$ORCID_NUMBER/person"
```

### Recipe for getting a list of ORCID Works 

This script will get an ORCID works list by ORCID number. It requires the following environment variable to be set.


+ ORCID_API_URL
+ ORCID_ACCESS_TOKEN
+ ORICD_NUMBER (can be provided as an option to the script)

```shell
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
        -X GET "$ORCID_API_URL/v2.0/$ORCID_NUMBER/works"
```


