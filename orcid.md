
# USAGE

    orcid [OPTIONS] ORCID

## SYSNOPIS

orcid is a command line tool for harvesting ORCID data from the ORCID API.
See http://orcid.org/organizations/integrators for details. It requires
a client id and secret to access. This is set via environment variables
or the command line.

## CONFIGURATION

+ ORCID_API_URL - set the URL for accessing the ORCID API (e.g. sandbox or members URL)
+ ORCID_CLIENT_ID - the client id for your registered ORCID app
+ ORCID_SECRET - the client secret needed to aquire an access token for the AP

## OPTIONS

```
	-activities	display activity
	-affiliations	display affiliations
	-bio	display biography
	-funding	display funding
	-h	display help
	-help	display help
	-l	display license
	-license	display license
	-o	use orcid id
	-orcid	use orcid id
	-profile	display profile
	-v	display version
	-version	display version
	-works	display works
```

## EXAMPLES

Get an ORCID "works" from the sandbox for a given ORCID id.

```
    export ORCID_API_URL="https://pub.sandbox.orcid.org"
	export ORCID_CLIENT_ID="APP-01XX65MXBF79VJGF"
	export ORCID_CLIENT_SECRET="3a87028d-c84c-4d5f-8ad5-38a93181c9e1"
	orcid -works 0000-0003-0900-6903
```

