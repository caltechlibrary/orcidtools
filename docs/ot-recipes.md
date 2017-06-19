
# OT Recipes

## Shell example to render works.bib

Goal is to use the "Works" ORCID API call to render a "works.bib" file.

### Requirements 

+ _orcid_ from [ot](https://caltechlibrary.github.io/ot)
+ _mkpage_ from [mkpage](https://caltechlibrary.github.io/mkpage)
+ _etc/orcid-api.bash_ a script holding your ORCID API credentials

### Example Shell commands

```shell
    . etc/orcid-api.bash
    orcid -works 0000-0003-0900-6903 > works.json
    mkpage "data=works.json" templates/works-to-bibtex.tmpl > works.bib
    cat works.bib
```
