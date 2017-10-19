
# orcidtools recipes


## Using the works end point to generate BibTeX

Goal of this recipe is to use the _orcid_ cli and the "-works-detailed" 
option to generate a BibTeX file of the works listed when you view your
ORCID profile.

### software used

+ _orcid_ from [orcidtools](https://caltechlibrary.github.io/orcidtools)
+ _mkpage_ from [mkpage](https://caltechlibrary.github.io/mkpage)

### configuration and templates

+ *etc/orcid-api.bash* a script holding your ORCID API credentials
    + see *etc/orcid-api.bash-example* in the orcidtools repository
+ *templates/works-detailed-to-bibtex.tmpl* a mkpage text template mapping the values to the desired bibtex output
    + from the orcidtools repository

### Configuration

The configuration needed by the _orcid_ cli is includes three environment variables
ORCID_API_URL, ORCID_CLIENT_ID and ORCID_CLIENT_SECRET. Here a sample of what an *etc/orcid-api.bash* 
might look like (note ORCID_CLIENT_ID and ORCID_CLIENT_SECRET are just placeholders as yours will
be different than mine).

```shell
    export ORCID_API_URL="https://pub.sandbox.orcid.org"
    export ORCID_CLIENT_ID="APP-ID-FROM-ORCID-DEVELOPMENT-TOOLS-GOES-HERE"
    export ORCID_CLIENT_SECRET="SOME_MIGHT_SECRET_THINGY3_GOES_HERE"
```

### Example Shell commands

```shell
    . etc/orcid-api.bash
    orcid -works-detailed 0000-0003-0900-6903 > works-detailed.json
    mkpage "data=works-detailed.json" templates/works-detailed-to-bibtex.tmpl > works-detailed.bib
    cat works-detailed.bib
```

The final *works-detailed.bib* file would look something like.

