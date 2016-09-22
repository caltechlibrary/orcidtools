#!/bin/bash

if [ ! -d "testresults" ]; then
    mkdir -p testresults
fi
mkpage "data=testdata/0000-0003-0900-6903.json" \
    templates/orcid2bibtex.tmpl 
