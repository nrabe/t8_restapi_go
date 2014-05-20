This is a personal project, comparing Python vs Go for API development.

The goal is to have equivalent projects in both Python and Go.

GO version: https://github.com/nrabe/t8_restapi_go

Python version: https://github.com/nrabe/t8_restapi_py


Go tools used:

Goclipse IDE, auto-formatting with CMD+SHIFT+F, (some) code completion and syntax check, and AppEngine friendly (sort of).
	https://code.google.com/p/goclipse/

Google Cloud SDK, to handle Appengine, Appengine GO SDK, Cloud SQL, Cloud Storage, etc. And the Google Eclipse plugin
	https://developers.google.com/cloud/sdk/
	https://developers.google.com/eclipse/docs/download

Gorilla Web Tookit to handle JSON-RPC, sessions, authentication and context... ah, and it's Google AppEngine friendly
	http://www.gorillatoolkit.org/

Google App Engine Datastore to store data.


API's implemented in both projects:

```
response = api.call('Region.Retrieve', {}) # response: {Count:0 Items:[]}

response = api.call('RestaurantDetail.Retrieve', {}) # response: {Count:0 Items:[]}

response = api.call('RestaurantDetail.Update', {Restaurant:{Uid: CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC Title: Regions:[] Tags:[] Details:}}) # response: {Count:0 Items:[]}

response = api.call('RestaurantTag.Retrieve', {}) # response: {Count:0 Items:[]}

response = api.call('System.CreateTestData', {GeneralArgs:{} CleanupOnly:false}) # response: {Warnings:[]}

response = api.call('System.Test', {Test:}) # response: {}
```

