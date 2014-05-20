#!/bin/bash

BASE_DIR="$( cd "$( dirname "$( dirname "${0}" )" )" && pwd )"

GOAPP=~/google-cloud-sdk/platform/google_appengine/goapp

GOPATH="$($GOAPP env GOPATH):$BASE_DIR" appcfg.py update -v .
