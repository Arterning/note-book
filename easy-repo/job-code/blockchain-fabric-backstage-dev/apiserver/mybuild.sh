#!/bin/bash
set -x
TARGET_PATH=$GOPATH/src/github.com/crc/zlzk/apiserver
CompileCmd='go build -tags "gm" --ldflags "-extldflags -static"'
BuildImage="crc/ubuntu:bcap-base"
docker run -it --rm -e GOPATH=$GOPATH -w ${TARGET_PATH} -v ${TARGET_PATH}:${TARGET_PATH} ${BuildImage} bash -c "${CompileCmd}"
cd -
