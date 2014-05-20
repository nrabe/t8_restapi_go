
Goclipse IDE, auto-formatting with CMD+SHIFT+F, (some) code completion and syntax check, and AppEngine friendly (sort of).
	https://code.google.com/p/goclipse/

Google Cloud SDK, to handle Appengine, Appengine GO SDK, Cloud SQL, Cloud Storage, etc. And the Google Eclipse plugin
	https://developers.google.com/cloud/sdk/
	https://developers.google.com/eclipse/docs/download

GORM as a db abstraction layer/orm... It has some minor things I dislike, but it seems pretty mature/well tested, and very well documented.
	https://github.com/jinzhu/gorm

	GORP seems nice enough, but besides the same things I dislike on GORM, it forces you to rely exclusively on writing SQL and concatenated string for SELECT ... WHERE and similar stuff
		https://github.com/coopernurse/gorp

Gorilla Web Tookit to handle JSON-RPC, sessions, authentication and context... ah, and it's Google AppEngine friendly
	http://www.gorillatoolkit.org/

JSON-RPC v2.0 ... nice and simple. Clearer, simpler and more suited to this than the RESTful meta-protocols. And it's Go-UnitTest friendly.
	http://www.jsonrpc.org/specification

Documentation with GODOC, the official documentation tool... not pretty, but std in GO world.
	http://localhost:6060/pkg/table8/restaurant_api/

Scripts to try and normalize the mess that GO + AppEngine + GoClipse paths and configurations are, to run several tools and perform ops.

Good sample of unit tests... both Go tests, and remote tests (remote tests: in a different language, even in another computer).


# MY SETUP of mysql:
####################

... interesting info about connections, pooling and what not: http://go-database-sql.org/accessing.html
	and https://groups.google.com/forum/#!topic/golang-nuts/je5WyDzib9I

Install MySQL Using MacPorts for MacOS X
	http://jackal.livejournal.com/2160464.html
	http://www.debuntu.org/how-to-create-a-mysql-database-and-set-privileges-to-a-user/

/opt/local/lib/mysql55/bin/mysql -u root -p
create database table8_gotest1;
grant usage on *.* to table8@localhost identified by 'table8';
grant all privileges on table8_gotest1.* to table8@localhost ;
grant usage on *.* to table8@127.0.0.1 identified by 'table8';
grant all privileges on table8_gotest1.* to table8@127.0.0.1 ;



# MY SETUP of dependencies:
###########################

# NOTE: sql drivers *MUST* be installed in the goapp path, not the local one, to avoid: panic: sql: Register called twice for driver postgres
goapp get github.com/go-sql-driver/mysql
goapp get github.com/lib/pq
GOPATH=/Users/nrabe/Documents/workspace/table8_restaurant_api/ goapp get github.com/go-sql-driver/mysql
#GOPATH=/Users/nrabe/Documents/workspace/table8_restaurant_api/ goapp get github.com/lib/pq
GOPATH=/Users/nrabe/Documents/workspace/table8_restaurant_api/ goapp get github.com/coopernurse/gorp
#GOPATH=/Users/nrabe/Documents/workspace/table8_restaurant_api/ goapp get github.com/jinzhu/gorm
GOPATH=/Users/nrabe/Documents/workspace/table8_restaurant_api/ goapp get github.com/gorilla/context
GOPATH=/Users/nrabe/Documents/workspace/table8_restaurant_api/ goapp get github.com/gorilla/sessions
GOPATH=/Users/nrabe/Documents/workspace/table8_restaurant_api/ goapp get github.com/gorilla/securecookie
GOPATH=/Users/nrabe/Documents/workspace/table8_restaurant_api/ goapp get github.com/gorilla/rpc
GOPATH=/Users/nrabe/Documents/workspace/table8_restaurant_api/ goapp get github.com/gorilla/rpc/v2/json

GOPATH=/Users/nrabe/Documents/workspace/table8_restaurant_api/ goapp get github.com/gorilla/context
GOPATH=/Users/nrabe/Documents/workspace/table8_restaurant_api/ goapp get github.com/gorilla/sessions
GOPATH=/Users/nrabe/Documents/workspace/table8_restaurant_api/ goapp get github.com/gorilla/securecookie
GOPATH=/Users/nrabe/Documents/workspace/table8_restaurant_api/ goapp get github.com/gorilla/rpc
GOPATH=/Users/nrabe/Documents/workspace/table8_restaurant_api/ goapp get github.com/gorilla/rpc/v2/json
GOPATH=/Users/nrabe/Documents/workspace/table8_restaurant_api/ goapp get github.com/gorilla/rpc/v2/json2

GOPATH=/Users/nrabe/Documents/workspace2/table8_restaurant_api/ goapp get github.com/gorilla/rpc/v2
GOPATH=/Users/nrabe/Documents/workspace2/table8_restaurant_api/ goapp get github.com/gorilla/rpc/v2/json2

