
# Demo using _orcid_ api tool 

## Known ORCID

+ Stephen Davison, [0000-0003-0102-8200](https://orcid.org/0000-0003-0102-8200)
+ Gail Clement, [0000-0001-5494-4806](https://orcid.org/0000-0001-5494-4806)
+ Tom Morrel, [0000-0001-9266-5146](https://orcid.org/0000-0001-9266-5146)
+ Robert Doiel, [0000-0003-0900-6903](https://orcid.org/0000-0003-0900-6903)


## Someone who will never have an ORCID

+ Richard Fenyman


## Tool Examples

### Setup

```shell
    #!/bin/bash
    #
    
    #
    # This is an example of the environment configuration for accessing the ORCID Public API sandbox.
    # You need to modify these values to support your access to the public API or sandbox.
    # These values were taken from the public documentation on the API
    #
    # See: https://members.orcid.org/api/accessing-public-api
    # See: http://members.orcid.org/api/tutorial-retrieve-data-public-api-curl-12-and-earlier
    #
    export ORCID_API_URL="https://pub.sandbox.orcid.org"
    export ORCID_CLIENT_ID="APP-01XX65MXBF79VJGF"
    export ORCID_CLIENT_SECRET="3a87028d-c84c-4d5f-8ad5-38a93181c9e1"
    # You will need to manually set ORCID_ACCESS_TOKEN based on the response from
    # the request to /oauth/token providing providing the client_id and client_secret.
    # Once you have the JSON blob back you can then set the ORCID_ACCESS_TOKEN variable
    # and export it into your environment.
    # E.g. `export ORCID_ACCESS_TOKEN="3b5bd7e6-8499-40ac-ac21-7301b09d4aab"`
```

Update the _etc/orcid-api.bash_ to your orcid api key and credentials. Then run --

```shell
    . etc/orcid-api.bash
```

### Search Examples

```shell
    orcid -search 'Stephen AND Davison AND (Caltech OR "California Institute of Technology")'
    orcid -search 'Gail AND Clement AND (Caltech OR "California Institute of Technology")'
    orcid -search 'Thomas AND Morrel AND (Caltech OR "California Institute of Technology")'
    orcid -search 'Robert AND Doiel AND (Caltech OR "California Institute of Technology")'
    orcid -search 'Richard AND Feynman AND (Caltech OR "California Institute of Technology")'
```

### Fetch ORCID records

```shell
    # Stephen Davison
    orcid -record 0000-0003-0102-8200
    # Gail Clement
    orcid -record 0000-0001-5494-4806
    # Tom Morrel
    orcid -record 0000-0001-9266-5146
    # Robert Doiel
    orcid -record 0000-0003-0900-6903
```
