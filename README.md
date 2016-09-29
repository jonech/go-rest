# go-rest
RESTful API with go and postgres


Prior to Execute
----------------
1. Set up a postgres database, create table like users.sql

2. Change the username, password, database and port in main.go

3. Get the 3rd party libraries/framework

External libraries installation
-------------------------------
```
$ go get github.com/lib/pq
$ go get github.com/gorilla/mux
```

Execute
--------
```
$ go build
$ ./go-rest
```

Usage
------
Find user
=========
```
http://localhost:<port>/user/:id
```
Create user
===========
```
$ curl -XPOST -H 'Content-Type: application/json' -d '{"name": "Bob Smith", "age": 50}' http://localhost:<port>/user
```

