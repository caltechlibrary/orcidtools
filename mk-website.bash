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

# Build utility docs pages
MakeSubPagesNav docs > docs/nav.md
MakeSubPagesIndex docs > docs/index.md
MakeSubPages docs

# Build how-to pages
#MakeSubPages how-to
