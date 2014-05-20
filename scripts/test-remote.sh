#!/bin/bash

BASE_DIR="$( cd "$( dirname "$( dirname "${0}" )" )" && pwd )"

GOAPP=~/google-cloud-sdk/platform/google_appengine/goapp

time env python tests_remote/test_errors.py
RV=$?; if [ $RV -ne 0 ]; then exit $RV; fi

time env python tests_remote/test_general.py
RV=$?; if [ $RV -ne 0 ]; then exit $RV; fi
