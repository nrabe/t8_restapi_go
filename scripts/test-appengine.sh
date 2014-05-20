#!/bin/bash

BASE_DIR="$( cd "$( dirname "$( dirname "${0}" )" )" && pwd )"

GOAPP=~/google-cloud-sdk/platform/google_appengine/goapp

ENDPOINT=http://t8-restapi-go.appspot.com/restaurant/0.1/

env python tests_remote/test_errors.py "$ENDPOINT"
RV=$?; if [ $RV -ne 0 ]; then exit $RV; fi

env python tests_remote/test_general.py "$ENDPOINT"
RV=$?; if [ $RV -ne 0 ]; then exit $RV; fi
