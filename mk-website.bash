#!/bin/bash
#

PROJECT=ot

function makePage () {
    title=$1
    page=$2
    nav=$3
    html_page=$4
    echo "Generating $html_page"
    mkpage \
        "title=text:$title" \
        "content=$page" \
        "nav=$nav" \
        page.tmpl > $html_page
}

if [ ! -f nav.md ]; then

    cat <<EOF> nav.md
+ [Home](/)
+ [README](index.html)
+ [Install](install.html)
+ [LICENSE](license.html)
+ [Github](https://github.com/caltechlibrary/$PROJECT)
EOF
    if [ -f NOTES.md ]; then 
        echo "+ [Notes](notes.html)" >> nav.md
    fi
fi


# index.html
if [ -f README.md ]; then
    makePage "$PROJECT" README.md nav.md index.html
fi

# install.html
if [ -f INSTALL.md ]; then
    makePage "$PROJECT" INSTALL.md nav.md install.html
fi

# license.html
if [ -f LICENSE ]; then
    makePage "$PROJECT" "markdown:$(cat LICENSE)" nav.md license.html
fi

# notes.html
if [ -f NOTES.md ]; then
    makePage "$PROJECT" NOTES.md nav.md notes.html
fi

# todo.html
if [ -f TODO.md ]; then
    makePage "$PROJECT" TODO.md nav.md todo.html
fi

# Add the files to git as needed
git add index.html install.html license.html notes.html todo.html
