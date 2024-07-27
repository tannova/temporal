# How to using demo temporal

## Temporal - Client
**This is service support execute workflow** 

`After run temporal , using cmd : go run cmd/main.go`

The system will run and register the API on port: 8083.
There are two ways to call the API:

* Method 1: Using curl:
` curl -X GET http://localhost:8080/weather\?city\=Cairo`


* Method 2: Using Postman:
`http://localhost:8083/weather?city=test`


## Temporal - Server
**This is service handle logic register workflow & acitviy and implement logic**

`After run temporal , using cmd : go run cmd/main.go
`
## How to start Temporal
There are two ways to start Temporal for the system.

Method 1: Using Docker Compose: 

- At the project location, run the command: docker-compose up

Method 2: Install Temporal on your machine

- Windows / MacOS: Install through the link: https://learn.temporal.io/getting_started/go/dev_environment/
Then run the command: temporal server start-dev

After starting Temporal:

`The server will run on port: 7233`

`Temporal UI port: 8233`

Refer to: https://learn.temporal.io/getting_started/go/dev_environment/
