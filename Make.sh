#!/usr/bin/env bash

CURRDIR=`pwd`
cd ../../../..
export GOPATH=`pwd`
cd ${CURRDIR}

go build -v -o tabtoy.exe github.com/vikingsc2007/tabtoy