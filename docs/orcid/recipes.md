
# orcidtools recipes

## Shell example to render works.bib

Goal is to use the "Works" ORCID API call to render a "works-detailed.bib" file.
Actually we want to use an _orcid_ cli option called "-works-detailed".
It calls the "Works" ORCID API end point to get the work ids, then 
gets the detailed works information for each.

### software used

+ _orcid_ from [orcidtools](https://caltechlibrary.github.io/orcidtools)
+ _mkpage_ from [mkpage](https://caltechlibrary.github.io/mkpage)

### configuration and templates

+ _etc/orcid-api.bash_ a script holding your ORCID API credentials
    + see _etc/orcid-api.bash-example_ in the orcidtools repository
+ _templates/works-detailed-to-bibtex.tmpl_ a mkpage text template mapping the values to the desired bibtex output
    + from the orcidtools repository

### Example Shell commands

```shell
    . etc/orcid-api.bash
    orcid -works-detailed 0000-0003-0900-6903 > works-detailed.json
    mkpage "data=works.json" templates/works-to-bibtex.tmpl > works-detailed.bib
    cat works-detailed.bib
```
