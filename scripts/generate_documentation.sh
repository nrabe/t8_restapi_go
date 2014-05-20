#!/bin/bash

BASE_DIR="$( cd "$( dirname "$( dirname "${0}" )" )" && pwd )"

GODOC=~/google-cloud-sdk/platform/google_appengine/godoc
GOAPP=~/google-cloud-sdk/platform/google_appengine/goapp

export GOPATH="$($GOAPP env GOPATH):$BASE_DIR" 
echo GOPATH=$GOPATH
$GODOC -v=true -http=:6060 &
sleep 2
cd doc && wget -r -nH -np -E -p -k 'http://localhost:6060/pkg/table8/restaurant_api/'

open pkg/table8/restaurant_api/index.html
trap "kill 0" SIGINT SIGTERM EXIT
