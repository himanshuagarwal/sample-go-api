# sample-go-api
A simple RESTful API (CRUD) built with GoLang

## Install
Clone and build :

```bash
$ git clone https://github.com/himanshuagarwal272/sample-go-api.git
$ cd sample-go-api
$ go build
```

## Running
First, setup docker-compose:

```bash
$ go run main.go
```

Then start querying at `http://localhost:3000/api/users/`

## APIs
The entity **User** has the following fields:

- ID (int)
- name (string)
- email (string)

Follows the list of users APIs:

|METHOD|URL|REQUEST HEADERS|REQUEST PAYLOAD|RESPONSE HEADERS|RESPONSE PAYLOAD|
|------|---|---------------|---------------|----------------|----------------|
|GET|http://localhost:3000/users/ | | | |User[]|
|POST|http://localhost:3000/user/10 |Content-Type: "application/json"|User||User|
|GET|http://localhost:3000/user/10 | | | |User|
|PUT|http://localhost:3000/user/10 |Content-Type: "application/json"|User||User|
|DELETE|http://localhost:3000/user/10 | | | | |			

