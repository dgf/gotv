#!/bin/sh
set -ev

# main package
main=github.com/dgf/gotv

# godoc.org refresh URL
godocURL=https://godoc.org/-/refresh

# refresh godoc
godocRefresh() {
	curl -D - -X POST -d "path=$1" $godocURL
}

# goreport.org refresh URL
goreportURL=https://goreportcard.com/checks

# refresh goreport
goreportRefresh() {
	curl -D - -X POST -d "repo=$1" $goreportURL
}

goreportRefresh "$main"
godocRefresh "$main"
for package in $(find . -mindepth 1 -type d | grep -v .git | grep -v testdata | cut -d '/'  -f 2); do
	godocRefresh "$main/$package"
done
