#!/bin/bash

function softwareCheck() {
	for CMD in "$@"; do
		APP=$(which "$CMD")
		if [ "$APP" = "" ]; then
			echo "Skipping, missing $CMD"
			exit 1
		fi
	done
}

function MakePage() {
	nav="$1"
	content="$2"
	html="$3"

	echo "Rendering $html"
	mkpage \
		"nav=$nav" \
		"content=$content" \
		page.tmpl >"$html"
	git add "$html"
}

function MakeSubPagesNav() {
    DIR="$1"
    START=$(pwd)
    cd "$DIR"
    echo "+ [Home](/)"
    echo "+ [Index](./)"
    echo "+ [Up](../)"
    for FNAME in $(ls *.md | sort); do
        if [ "$FNAME" != "nav.md" ] && [ "$FNAME" != "index.md" ]; then
            HNAME=$(basename $FNAME ".md")
            TITLE=$(titleline -i "$FNAME")
            echo "+ [$TITLE](${HNAME}.html)"
        fi
    done
    cd "$START"
}

function MakeSubPagesIndex() {
    DIR="$1"
    START=$(pwd)
    cd "$DIR"
    echo ""
    echo "# Documentation"
    echo ""
    for FNAME in $(ls *.md | sort); do
        if [ "$FNAME" != "nav.md" ] && [ "$FNAME" != "index.md" ]; then
            HNAME=$(basename $FNAME ".md")
            TITLE=$(titleline -i "$FNAME")
            echo "+ [$TITLE](${HNAME}.html)"
        fi
    done
    cd "$START"
}

function MakeSubPages() {
    SUBDIR="${1}"
    find "${SUBDIR}" -type f | grep -E '\.md$' | while read FNAME; do
        FNAME="$(basename "${FNAME}" ".md")"
        if [ "$FNAME" != "nav" ]; then
	        MakePage "${SUBDIR}/nav.md" "${SUBDIR}/${FNAME}.md" "${SUBDIR}/${FNAME}.html"
        fi
    done
}

echo "Checking software..."
softwareCheck mkpage
echo "Generating website"
MakePage nav.md README.md index.html
MakePage nav.md INSTALL.md install.html
MakePage nav.md "markdown:$(cat LICENSE)" license.html
MakePage docs/nav.md docs/index.md docs/index.html
MakePage docs/nav.md docs/reference.md docs/reference.html
MakePage docs/orcid/nav.md docs/orcid/index.md docs/orcid/index.html
MakePage docs/orcid/nav.md docs/orcid/v2.0_API_end_points.md docs/orcid/v2.0_API_end_points.html
MakePage docs/orcid/nav.md docs/orcid/recipes.md docs/orcid/recipes.html
MakePage docs/orcid/nav.md docs/orcid/demo.md docs/orcid/demo.html


