
# ot

  Orcid Tools

Library for making command line tools that interact with the ORCID Public API

## Configuration

The commands available with the _ot_ package require three environment variables to be defined.

+ ORCID_API_URL
+ ORCID_CLIENT_ID
+ ORCID_CLIENT_SECRET

If you want to use the API URL https://pub.orcid.org then you'll seen to register your application
to generate your client id and secret.  If you want to experiment with the orcid public api to
test code (e.g. say test this package) you can use the example client id and secret describe on the
orcid.org [website](http://members.orcid.org/api/tutorial-retrieve-data-public-api-curl-12-and-earlier)
along with the API URL https://pub.sandbox.orcid.org.

Below is an example of using the sandbox credentials and to get an orcid bio.

```shell
    export ORCID_API_URL="https://pub.sandbox.orcid.rg"
    export ORCID_CLIENT_ID="APP-01XX65MXBF79VJGF"
    export ORCID_CLIENT_SECRET="3a87028d-c84c-4d5f-8ad5-38a93181c9e1""
    go run cmds/orcidapi/orcidapi.go '{"path":"/orcid-bio","orcid":"0000-0002-2389-8429"}'
```

Should return a response like

```json
    {
        "message-version":"1.2",
        "orcid-profile":{
            "orcid":null,
            "orcid-id":null,
            "orcid-identifier":{
                "value":null,
                "uri":"http://sandbox.orcid.org/0000-0002-2389-8429",
                "path":"0000-0002-2389-8429",
                "host":"sandbox.orcid.org"
            },
            "orcid-deprecated":null,
            "orcid-preferences":{"locale":"EN"},
            "orcid-history":{
                "creation-method":"DIRECT",
                "completion-date":null,
                "submission-date":{"value":1414132840517},
                "last-modified-date":{"value":1461250703013},
                "claimed":{"value":true},
                "source":null,
                "deactivation-date":null,
                "verified-email":{"value":true},
                "verified-primary-email":{"value":true},
                "visibility":null
            },
            "orcid-bio":{
                "personal-details":{
                    "given-names":{"value":"Sofia","visibility":null},
                    "family-name":{"value":"Hernandez","visibility":null},
                    "credit-name":null,
                    "other-names":null
                },
                "biography":null,
                "researcher-urls":null,
                "contact-details":null,
                "keywords":null,
                "external-identifiers":{
                    "external-identifier":[
                        {
                            "orcid":null,
                            "external-id-orcid":null,
                            "external-id-common-name":{"value":"Loop profile"},
                            "external-id-reference":{"value":"559"},
                            "external-id-url":{"value":"http://loop.frontiers-sandbox-int.info/people/559/overview?referrer=orcid_profile"},
                            "external-id-source":null,
                            "source":{
                                "source-orcid":null,
                                "source-client-id":{
                                    "value":null,
                                    "uri":"http://sandbox.orcid.org/client/APP-674MCQQR985VZZQ2",
                                    "path":"APP-674MCQQR985VZZQ2","host":"sandbox.orcid.org"
                                },
                                "source-name":{"value":"Display \"name\" for testing"},
                                "source-date":{"value":1461170594042}
                            }
                        },{
                            "orcid":null,
                            "external-id-orcid":null,
                            "external-id-common-name":{"value":"put public 1.2 2"},
                            "external-id-reference":{"value":"15"},
                            "external-id-url":{"value":"www.myid.com/4567888"},
                            "external-id-source":null,
                            "source":{
                                "source-orcid":{
                                    "value":null,
                                    "uri":"http://sandbox.orcid.org/0000-0003-2736-806X",
                                    "path":"0000-0003-2736-806X",
                                    "host":"sandbox.orcid.org"
                                },
                                "source-client-id":null,
                                "source-name":{"value":"Name of your client application"},
                                "source-date":{"value":1461250665148}
                            }
                        },{
                            "orcid":null,
                            "external-id-orcid":null,
                            "external-id-common-name":{"value":"put public 1.2 3"},
                            "external-id-reference":{"value":"16"},
                            "external-id-url":{"value":"www.myid.com/45678888"},
                            "external-id-source":null,
                            "source":{
                                "source-orcid":{
                                    "value":null,
                                    "uri":"http://sandbox.orcid.org/0000-0003-2736-806X",
                                    "path":"0000-0003-2736-806X",
                                    "host":"sandbox.orcid.org"
                                },
                                "source-client-id":null,
                                "source-name":{"value":"Name of your client application"},
                                "source-date":{"value":1461250703010}
                            }
                        }
                    ],
                    "visibility":"PUBLIC"
                },
                "delegation":null,
                "scope":null
            },
            "orcid-activities":null,
            "orcid-internal":null,
            "type":"USER",
            "group-type":null,
            "client-type":null
        },
        "orcid-search-results":null,
        "error-desc":null
    }
```

## Reference

+ [orcid.org](http://orcid.org)
+ [ORCID Public API Documentation](http://members.orcid.org/api/introduction-orcid-public-api)
+ [Tutorial on getting ORCID with CURL](http://members.orcid.org/api/tutorial-retrieve-orcid-id-curl-v12-and-earlier)
+ [Code Examples](http://members.orcid.org/api/code-examples)
+ [Working with GZip and Tar](http://blog.ralch.com/tutorial/golang-working-with-tar-and-gzip/) - golang example useful for parsing the tar ball of publicly released data
