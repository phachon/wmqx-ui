#!/bin/bash
VER=$1
if [ "$VER" = "" ]; then
    echo 'please input pack version!'
    exit 1
fi
RELEASE="release-${VER}"
rm -rf release-*
mkdir ${RELEASE}

# windows amd64
echo 'Start pack windows amd64...'
GOOS=windows GOARCH=amd64 go get ./...
GOOS=windows GOARCH=amd64 go build ./
tar -czvf "${RELEASE}/wmqx-ui-windows-amd64.tar.gz" wmqx-ui.exe conf/default.conf db/ logs/.gitignore static/ views/ LICENSE README.md
rm -rf wmqx-ui.exe

echo 'Start pack windows X386...'
GOOS=windows GOARCH=386 go get ./...
GOOS=windows GOARCH=386 go build ./
tar -czvf "${RELEASE}/wmqx-ui-windows-386.tar.gz" wmqx-ui.exe conf/default.conf db/ logs/.gitignore static/ views/ LICENSE README.md
rm -rf wmqx-ui.exe

echo 'Start pack linux amd64'
GOOS=linux GOARCH=amd64 go get ./...
GOOS=linux GOARCH=amd64 go build ./
tar -czvf "${RELEASE}/wmqx-ui-linux-amd64.tar.gz" wmqx-ui conf/default.conf db/ logs/.gitignore static/ views/ LICENSE README.md
rm -rf wmqx-ui

echo 'Start pack linux 386'
GOOS=linux GOARCH=386 go get ./...
GOOS=linux GOARCH=386 go build ./
tar -czvf "${RELEASE}/wmqx-ui-linux-386.tar.gz" wmqx-ui conf/default.conf db/ logs/.gitignore static/ views/ LICENSE README.md
rm -rf wmqx-ui

echo 'Start pack mac amd64'
GOOS=darwin GOARCH=amd64 go get ./...
GOOS=darwin GOARCH=amd64 go build ./
tar -czvf "${RELEASE}/wmqx-ui-mac-amd64.tar.gz" wmqx-ui conf/default.conf db/ logs/.gitignore static/ views/ LICENSE README.md
rm -rf wmqx-ui

echo 'END'
