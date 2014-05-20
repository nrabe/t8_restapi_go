#!/usr/bin/env python
import sys
import urllib2
import json
from _test import *
import time

endpoint = SERVICE_ENDPOINT
if len(sys.argv) > 1:
    endpoint = sys.argv[1]

api = SimpleJsonRPCClient(endpoint)


# HTTP GET call is invalid (in fact, ANYTHING but HTTP POST )
try:
    req = urllib2.Request(SERVICE_ENDPOINT, None, {'Content-Type': 'application/json'})
    f = urllib2.urlopen(req)
    response = json.loads(f.read())
    f.close()
    if response == '':
        raise urllib2.HTTPError(500, 'response expected')
except urllib2.HTTPError, e:
    assert e.code in (405, 500), 'Unexpected error: %r %r %r' % (e.code, e.reason, e.read())
    print 'OK. It was an expected error HTTP %r' % e.code


# HTTP POST with incorrect JSON
try:
    req = urllib2.Request(SERVICE_ENDPOINT, 'xxx', {'Content-Type': 'application/json'})
    f = urllib2.urlopen(req)
    response = f.read()
    if response == '':
        raise urllib2.HTTPError(SERVICE_ENDPOINT, 500, 'response expected', {}, None)
except urllib2.HTTPError, e:
    assert e.code in (405, 500), 'Unexpected error: %r %r %r' % (e.code, e.reason, e.read())
    print 'OK. It was an expected error HTTP %r' % e.code


# HTTP POST with empty string
try:
    req = urllib2.Request(SERVICE_ENDPOINT, '', {'Content-Type': 'application/json'})
    f = urllib2.urlopen(req)
    response = f.read()
    if response == '':
        raise urllib2.HTTPError(SERVICE_ENDPOINT, 500, 'response expected', {}, None)
except urllib2.HTTPError, e:
    assert e.code in (405, 500), 'Unexpected error: %r %r %r' % (e.code, e.reason, e.read())
    print 'OK. It was an expected error HTTP %r' % e.code


# HTTP POST with empty JSON
try:
    req = urllib2.Request(SERVICE_ENDPOINT, json.dumps({}), {'Content-Type': 'application/json'})
    f = urllib2.urlopen(req)
    response = f.read()
except urllib2.HTTPError, e:
    assert e.code in (405, 500), 'Unexpected error: %r %r %r' % (e.code, e.reason, e.read())
    print 'OK. It was an expected error HTTP %r' % e.code


# HTTP GET call is invalid (in fact, ANYTHING but HTTP POST )
try:
    req = urllib2.Request(SERVICE_ENDPOINT, None, {'Content-Type': 'dummy/json'})
    f = urllib2.urlopen(req)
    response = json.loads(f.read())
    f.close()
except urllib2.HTTPError, e:
    assert e.code in (405, 500), 'Unexpected error: %r %r %r' % (e.code, e.reason, e.read())
    print 'OK. It was an expected error HTTP %r' % e.code

# service does not exists
try:
    response = api.call('Dummy-Doesnot.Retrieve', Id=1)
    assert False, 'Expecting an error'
except SimpleJsonRPCClientError, e:
    #assert '-32601' in repr(e), 'Unexpected error: %r' % e
    print 'OK. It was an expected error %r' % e.code

# incorrect parameter type
try:
    response = api.call('System.Test', test=1)
    assert False, 'Expecting an error'
except SimpleJsonRPCClientError, e:
    assert e.code == 500, 'Unexpected error: %r' % e
    print 'OK. It was an expected error %r' % e.code

# incorrect parameter name
try:
    response = api.call('System.Test', test_this_param_does_not_exist=1)
    assert False, 'Expecting an error'
except SimpleJsonRPCClientError, e:
    assert e.code == 500, 'Unexpected error: %r' % e
    print 'OK. It was an expected error %r' % e.code

# missing a required parameter
try:
    response = api.call('System.Test')
    assert False, 'Expecting an error'
except SimpleJsonRPCClientError, e:
    #assert 'InvalidParamsError' in repr(e), 'Unexpected error: %r' % e
    print 'OK. It was an expected error %r' % e.code

# programming error on the server side
try:
    response = api.call('System.Test', Test="fatal")
    assert False, 'Expecting an error'
except SimpleJsonRPCClientError, e:
    assert e.code in (405, 500), 'Unexpected error: %r %r %r' % (e.code, e.reason, e.read())
    print 'OK. It was an expected error HTTP %r' % e.code

response = api.call('System.Test', Test='success')
print 'success System.Test=', response
