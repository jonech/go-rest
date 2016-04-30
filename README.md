# travel-bug-backend
Backend for android app


Prior to Execute
----------------
1. Set up a postgres database on your local machine

2. Change the username, password, database in main.go

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
$ ./travel-bug-backend
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


Useful Tutorial
---------------
* [create a restful api](http://thenewstack.io/make-a-restful-json-api-go/)
* [create restful api + model & controller seperation + mongo](http://stevenwhite.com/building-a-rest-service-with-golang-2/)
* [how to postgres with golang](https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/05.4.html)
