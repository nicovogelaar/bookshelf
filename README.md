# Bookshelf [![Build Status](https://travis-ci.org/nicovogelaar/bookshelf.svg?branch=master)](https://travis-ci.org/nicovogelaar/bookshelf)

This is a simple bookshelf application written in golang.

There are two parts of the application:

1. **api** - this is the server side app with various api endpoints. The api is writen in golang.
2. **web** - this is the frontend app that is consuming the api endpoints. The web app is created with Vue.js.

## Running the apps

1. Run the API
```
$ make
```

2. Run the web application
```
$ cd ./web
$ make
```

Once the docker images are successfully built and the containers are up, the application should be available at the following address:

```
http://127.0.0.1:3000
```

## Running the tests

```
$ make test
```
