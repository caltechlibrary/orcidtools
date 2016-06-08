
+ [Public API Docs](http://members.orcid.org/api/introduction-orcid-public-api)

# Get All Caltech ORCID ids

+ [Getting ORCIDs by institution](http://members.orcid.org/finding-orcid-record-holders-your-institution) - get orcids with @caltech.edu in the name
+ [Tutorial on Search for ORCIDs](http://members.orcid.org/api/tutorial-searching-api-12-and-earlier) - 
+ [Tutorial on searching the API with Curl](http://members.orcid.org/api/tutorial-retrieve-data-public-api-curl-12-and-earlier)
+ [How do I get public data](http://support.orcid.org/knowledgebase/articles/223698)

# Steps 

1. Get an Access token (does not expire)

```shell
    curl -i -L -H "Accept: application/json" -d "client_id=APP-01XX65MXBF79VJGF" -d "client_secret=3a87028d-c84c-4d5f-8ad5-38a93181c9e1" -d "scope=/read-public" -d "grant_type=client_credentials" "https://pub.sandbox.orcid.org/oauth/token"
```

```json

    {"access_token":"39a1d9c5-e753-41b8-b676-a82142ef67ae","token_type":"bearer","refresh_token":"a3a2420a-b964-4cee-bc15-7bdd873b1643","expires_in":631138518,"scope":"/read-public","orcid":null}
```

2. Using access tokens search by email domain

```shell
curl -H "Content-Type: application/orcid+xml" -H "Authorization: 39a1d9c5-e753-41b8-b676-a82142ef67ae" "https://pub.sandbox.orcid.org/v1.2/search/orcid-bio/?q=email:*@caltech.edu
```


3. Save data 
4. transform into XLSX/CSV files

## Algorithm

1. For each faculty name in faculty list matching eprints
2. For each DOI associated with an EPrint query the Orcid Public API for associated Name and DOI, if hit remember the information and ORCID
3. For each faculty member without a related DOI query the Orcid Public API by name, if only one record return remember the orcid and return any associated data
4. For each EPrint with ORCID make sure we've have added the name to faculty list, to get full details query the Orcid Public API and save record

