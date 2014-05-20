# /usr/bin/env python
import urllib2
import json
from _test import *
import time

api = SimpleJsonRPCClient(SERVICE_ENDPOINT)

response = api.call('System.CreateTestData')
time.sleep(0.5)  # I *HATE* AppEngine DataStore "eventual consistency"

response = api.call('Region.Retrieve')

response = api.call('RestaurantTag.Retrieve')

response = api.call('RestaurantDetail.Retrieve')
assert response['Items'][0]['Title'] == 'TEST McDonalds'
restaurant = response['Items'][0]

response = api.call('RestaurantDetail.Update', Title=restaurant['Title'] + ' #1', Tags=['Asian fusion'], Details='')
assert response['Items'][0]['Title'] == 'TEST McDonalds #1', response['Items'][0]['Title']
assert response['Items'][0]['Tags'] == ['Asian fusion'], response['Items'][0]['Tags']
assert response['Items'][0]['Details'] == '', response['Items'][0]['Details']

response = api.call('System.CreateTestData', CleanupOnly=True)
