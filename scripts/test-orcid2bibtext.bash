#!/bin/bash

if [ ! -d "testresults" ]; then
    mkdir -p testresults
fi
for ORCID_ID in "0000-0003-0900-6903" "0000-0003-0248-0813" "0000-0001-5494-4806"; do
    # DEBUG
    scripts/api-get-profile.sh "$ORCID_ID" > "testdata/$ORCID_ID.json"
    #mkpage "data=testdata/$ORCID_ID.json" \
    #    templates/orcid2bibtex.tmpl > testresults/$ORCID_ID.bib
    # DEBUGGING
    #cat testresults/$ORCID_ID.bib
done

