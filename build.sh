#!/bin/sh

ver=$(git describe --tags 2> /dev/null)
if [ "$ver" == "" ]; then ver='master'; fi

go build -ldflags "-X main.version=$ver"