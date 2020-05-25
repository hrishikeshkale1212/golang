# flexiloans http-go-server

This repo contains a simple/basic HTTP server in Go, with a basic code organization.
We use:
* net/http package to start and serve HTTP server
* Gorilla mux to handle routes
* Swagger in lorder to serve a REST API compliant with OpenAPI specs

## Pre-requisits

Install Go in 1.13 version minimum.

## Build the app

`$ go build -o bin/http-go-server internal/main.go`

or

`$ make build`

## Run the app

`$ ./bin/http-go-server`

## Test the app

```
$ curl http://localhost:8080/kafka
OK

$ curl http://localhost:8080/kafka/publisher

```

|                 URL					 | Port | HTTP Method			       | Operation														    |
|:-------------------------:|:--------:|:-----------------------:|------------------------------------------------------------------------|
| /kafka							 | 8080 | GET       |  Test if the app is running							    |
| /kafka/publisher							 | 8080 | GET       |  push data to kafka							    |						    |


`$ curl localhost:8080/kafka/publisher`
