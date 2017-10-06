
# USAGE

## orcid [OPTIONS] ORCID

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
	-O	use orcid id
	-activities	display activities
	-address	display address
	-educations	display education affiliations
	-email	display email
	-employments	display employment affiliations
	-external-ids	display external identifies
	-fundings	display funding activities
	-h	display help
	-help	display help
	-keywords	display keywords
	-l	display license
	-license	display license
	-orcid	use orcid id
	-other-names	display other names
	-peer-reviews	display peer review activities
	-person	display person
	-personal-details	display personal detials
	-record	display record
	-researcher-urls	display researcher urls
	-search	search for terms
	-v	display version
	-verbose	enable verbose logging
	-version	display version
	-works	display 
```


## EXAMPLES

Get an ORCID "works" from the sandbox for a given ORCID id.

```
    export ORCID_API_URL="https://pub.sandbox.orcid.org"
    export ORCID_CLIENT_ID="APP-01XX65MXBF79VJGF"
    export ORCID_CLIENT_SECRET="3a87028d-c84c-4d5f-8ad5-38a93181c9e1"
    orcid -works 0000-0003-0900-6903
```


orcid v0.0.4
