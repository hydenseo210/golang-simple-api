# hyden-simple-api
Just a simple api

## Application Details
Application implements the following:

a simple root endpoint which responds in a simple manner; "hello world".
a health endpoint which returns an appropriate response code
a metadata endpoint which returns basic information about your application, including version and associated git hash: e.g.
```
"myapplication": [
  {
    "version": "1.0",
    "description" : "pre-interview technical test",
    "lastcommitsha": "abc57858585"
  }
]
```
Curl commands
```
curl -I localhost:8080/health
curl localhost:8080/hello
curl localhost:8080/data
```
