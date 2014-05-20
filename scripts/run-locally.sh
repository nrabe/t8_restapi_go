#!/bin/bash

BASE_DIR="$( cd "$( dirname "$( dirname "${0}" )" )" && pwd )"

GOAPP=~/google-cloud-sdk/platform/google_appengine/goapp
#GOAPP=~/go_appengine/goapp
$GOAPP env
#GOPATH="$($GOAPP env GOPATH):$BASE_DIR" $GOAPP serve

GOPATH="$($GOAPP env GOPATH):$BASE_DIR" dev_appserver.py .
