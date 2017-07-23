#!/bin/sh
set -ev

# main package
main=github.com/dgf/gotv

# godoc.org refresh URL
godocURL=http://godoc.org/-/refresh

# refresh godoc
godocRefresh() {
	curl -D - -X POST -d "path=$1" $godocURL
}

godocRefresh "$main"
for package in $(find . -mindepth 1 -type d | grep -v .git | grep -v testdata | cut -d '/'  -f 2); do
	godocRefresh "$main/$package"
done
