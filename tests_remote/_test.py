# /usr/bin/env python
import urllib2
import json

SERVICE_ENDPOINT = 'http://localhost:8080/restaurant/0.1/'


class SimpleJsonRPCClientError(Exception):
    def __init__(self, code=500, message='no detail provided'):
        self.code = code
        self.message = message

    def __str__(self):
        return '!SimpleJsonRPCClientError: code=%s message=%r' % (self.code, self.message)

    def __unicode__(self):
        return '!SimpleJsonRPCClientError: code=%s message=%r' % (self.code, self.message)

    def __repr__(self):
        return '!SimpleJsonRPCClientError: code=%s message=%r' % (self.code, self.message)


class SimpleJsonRPCClient:
    """ very simple json rpc v2 client... suports single API calls, and batches of them """
    def __init__(self, endpoint):
        self.endpoint = endpoint
        self.id_counter = 1

    def callBatch(self, batch_calls):
        for call in batch_calls:
            call['jsonrpc'] = "2.0"
            call['id'] = self.id_counter
            self.id_counter += 1
        if len(batch_calls) == 1:
            batch_calls = batch_calls[0]
        data = json.dumps(batch_calls)
        print 'CALL.REQUEST %r' % data
        try:
            req = urllib2.Request(self.endpoint, data, {'Content-Type': 'application/json'})
            f = urllib2.urlopen(req)
            buf = f.read()
            if not buf:
                raise SimpleJsonRPCClientError(code=500, message='Empty response')
            response = json.loads(buf)
            f.close()
        except urllib2.HTTPError, e:
            raise SimpleJsonRPCClientError(code=e.code, message=e.reason)
        print 'CALL.RESPONS %r' % response
        return response

    def call(self, method, _params=None, **kwargs):
        _params = _params or {}
        _params.update(kwargs)
        response = self.callBatch([{'method': method, 'params': _params}])
        if 'error' in response:
            raise SimpleJsonRPCClientError(code=500, message=response['error'])
        return response['result']
