[![Go Report Card](http://goreportcard.com/badge/caltechlibrary/ot)](http://goreportcard.com/report/caltechlibrary/ot)

# ot

  Orcid Tools

A set of Bash scripts, Go template for working with the Public ORCID API.

## Configuration

Running the Bash scripts or command line tools built on the _ot_ package require specific environment variables to be defined.

+ ORCID_API_URL
+ ORCID_CLIENT_ID
+ ORCID_CLIENT_SECRET


If you want to use the API URL https://pub.orcid.org then you'll need to register an application
to generate your client id and secret.  If you want to experiment with the orcid public api to
test code (e.g. say test this package) you can use the example client id and secret describe on the
orcid.org [website](http://members.orcid.org/api/tutorial-retrieve-data-public-api-curl-12-and-earlier)
along with the API URL https://pub.sandbox.orcid.org.

The bash scripts provided in the repository rely on a few environment variables.
You can define those variables in a Bash script, sourcing that script will then
expose those variables in your current Bash session.

Below is an example of setup script that would be sourced to access the sandbox 

```shell
    #!/bin/bash
    export ORCID_API_URL="https://pub.sandbox.orcid.rg"
    export ORCID_CLIENT_ID="APP-01XX65MXBF79VJGF"
    export ORCID_CLIENT_SECRET="3a87028d-c84c-4d5f-8ad5-38a93181c9e1""
```

Assuming you saved this script as "etc/sandbox.sh" you would source it with the command

```shell
    . etc/sandbox.sh
```

You could then login to the API with

```shell
    ./scripts/api-login.sh
```

This will provide you with an Access token (you would cut and paste from the console to set that
into the environment). Once *ORCID_ACCESS_TOKEN* is defined in your environment you then can use
the other scripts to query the ORCID API for profile, bio and works data.

Putting it together

```shell
    . etc/sandbox.sh
    ./scripts/api-login.sh
    # Cut and past the 'export ORCID_ACCESS_TOKEN' line into the console
    # Then you can get the "works" for 0000-0003-0900-6903 with
    ./scripts/api-get-works.sh 0000-0003-0900-6903
```

## Reference

+ [orcid.org](http://orcid.org)
+ [ORCID Public API Documentation](http://members.orcid.org/api/introduction-orcid-public-api)
+ [Tutorial on getting ORCID with CURL](http://members.orcid.org/api/tutorial-retrieve-orcid-id-curl-v12-and-earlier)
+ [Code Examples](http://members.orcid.org/api/code-examples)
+ [Working with GZip and Tar](http://blog.ralch.com/tutorial/golang-working-with-tar-and-gzip/) - golang example useful for parsing the tar ball of publicly released data
+ [Useful ORCID API end points](http://members.orcid.org/api/tutorial-searching-api-12-and-earlier)

