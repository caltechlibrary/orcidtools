
# ot

  Orcid Tools

A command line tool called _orcid_, a set of Bash scripts and Go template for working with the Public ORCID API.

## Configuration

The _orcid_ tool and Bash scripts share a common configuration. These are set via environment variables.
The following are supported, the first three required.

+ ORCID_API_URL
+ ORCID_CLIENT_ID
+ ORCID_CLIENT_SECRET
+ ORCID_ACCESS_TOKEN (known on successful login to the API)


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

## the ORCID tool

The command line tool works simularly to the bash scripts. You source a configuration then run the tool. Unlike
the shell scripts login is automatic so you can focus on the command you need. The command line tool expacts 
an ORCID id as a command line parameter so it can get back a specific record.

```shell
    . etc/sandbox.sh
    orcid -works 0000-0003-0900-6903
```

Would list the works for the ORCID id of "0000-0003-0900-6903". The resulting document would be in JSON form.

Taking things a step further you can generate a BibTeX from the works in your ORCID using the _orcid_ tool and
[mkpage](https://caltechlibrary.github.io/mkpage) tool together with the templates included in this repository.

```shell
    . etc/sandbox.sh
    orcid -works 0000-0003-0900-6903 > 0000-0003-0900-6903-works.json
    mkpage "data=0000-0003-0900-6903-works.json" templates/orcid2bibtex.tmpl > 0000-0003-0900-6903.bib
```


## Working with the scripts

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
+ [Useful ORCID API end points](http://members.orcid.org/api/tutorial-searching-api-12-and-earlier)

