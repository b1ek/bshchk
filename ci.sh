#!/bin/sh

rm -f $(find . -type f | grep -E ^\./bshchk.\*)

GOOS=linux GOARCH=386 go build -o "bshchk.linux.i386" -ldflags "-s -w -X main.version=$VERSION" . &
GOOS=linux GOARCH=amd64 go build -o "bshchk.linux.amd64" -ldflags "-s -w -X main.version=$VERSION" . &
GOOS=linux GOARCH=arm64 go build -o "bshchk.linux.arm64" -ldflags "-s -w -X main.version=$VERSION" . &
GOOS=linux GOARCH=arm go build -o "bshchk.linux.arm" -ldflags "-s -w -X main.version=$VERSION" . &

GOOS=windows GOARCH=386 go build -o "bshchk.windows.i386.exe" -ldflags "-s -w -X main.version=$VERSION" . &
GOOS=windows GOARCH=amd64 go build -o "bshchk.windows.amd64.exe" -ldflags "-s -w -X main.version=$VERSION" . &
GOOS=windows GOARCH=arm64 go build -o "bshchk.windows.arm64.exe" -ldflags "-s -w -X main.version=$VERSION" . &
GOOS=windows GOARCH=arm go build -o "bshchk.windows.arm" -ldflags "-s -w -X main.version=$VERSION" . &

GOOS=darwin GOARCH=amd64 go build -o "bshchk.darwin.amd64" -ldflags "-s -w -X main.version=$VERSION" . &
GOOS=darwin GOARCH=arm64 go build -o "bshchk.darwin.arm64" -ldflags "-s -w -X main.version=$VERSION" . &

wait
