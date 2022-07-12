# Implementing JWT based Authentication for REST API in Golang with Fiber and jwt-go

## Introduction

Authentication is one of the common things to be considered
in building an API. Here is a simple code implementing JWT
application in a webserver built with the help of fiber 
framework.

## Installation

Go `1.18` version is used for this demo.

Clone the current repository with 
```bash
git clone github.com/smc181002/jwt-with-go-fiber.git
```

### Fiber CLI

A dev server can be started for Fiber using this CLI with 
live reload.

```bash
go install github.com/gofiber/cli/fiber
```

### Package installation

Install the packages used with the commands below
```bash
  go get github.com/gofiber/fiber/v2
  go get github.com/golang-jwt/jwt/v4
```

## Testing application

The user name is fixed as for this demo but in an actual 
application, the username will be fetched from database and
the username may be replaced or combined with other 
parameters like roles

### Using cURL

cURL is a tool installed in linux distributions by default 
used to fetch data from an endpoint.

**Get JWT token**

we can get the JWT token by POSTing to `/api/auth` endpoint

```bash
curl \
--data '{"name": "meher", "password": "passwd@123"}' \
-H 'Content-Type: application/json' \
http://localhost:3000/api/auth
```

save the output value in environmental variables from the 
response JSON in the above request which will be 
something like below

```json
{"token":"eyJh...XVCJ9.eyJ1...4ODQ0NH0.BZZltUp...2-Urx0HUfb-I"}
```

To request public endpoint, we can GETting `/api/posts` 
endpoint.

```bash
curl http://localhost:3000/api/posts/
```

To get the protected routes, we can send the token through 
Authorization Header as Bearer token

```bash
curl \
-H "Authorization: Bearer $TOKEN" \
http://localhost:3000/api/posts/private
```

