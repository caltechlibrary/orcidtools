
This is a placeholder for recipes using the _orcid_ tool that comes with the *ot* package
along with other tools developed at Caltech Library like [mkpage](https://caltechlibrary.github.io/mkpage).


NOTE: I need to develop an example along the lines of...

```shell
    . etc/mysetup.bash
    orcid -works 0000-0003-0900-6903 > works.json
    mkpage "data=works.json" templates/works-to-bibtex.tmpl > works.bib
    cat works.bib
```
