#!/bin/bash

BASE_DIR="$( cd "$( dirname "$( dirname "${0}" )" )" && pwd )"

GODOC=~/google-cloud-sdk/platform/google_appengine/godoc
GOAPP=~/google-cloud-sdk/platform/google_appengine/goapp

export GOPATH="$($GOAPP env GOPATH):$BASE_DIR" 
echo GOPATH=$GOPATH
$GODOC -v=true -http=:6060
