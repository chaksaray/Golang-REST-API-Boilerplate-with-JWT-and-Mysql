# Golang-REST-API-Boilerplate-with-JWT-and-Mysql

Golang RESTful API Boilerplate with JWT Authentication and backend Mysql. It covers the basic needs, and boilerplate work of a new project. It promotes the best practices that follow the clean architecture.

The Rest API provides the following features right out of the box:

-   Endpoints in the widely accepted format
-   Standard CRUD operations
-   JWT-based authentication
-   Middleware
-   Environment dependent application configuration
-   Logging
-   Error handling
-   Database migration and seeding
-   Data validation
-   Full test cover
-   Cache(Redis) integration
-   Docker compose
-   API doc using swagger as yaml file

It uses the following Go packages

-   Routing: [gorilla/mux](github.com/gorilla/mux)
-   Database access, migration and seeding: [jinzhu/gorm](github.com/jinzhu/gorm)
-   Env controll: [godotenv](github.com/joho/godotenv)
-   JWT: [jwt-go](github.com/dgrijalva/jwt-go)
-   Redis cache [go-redis](github.com/go-redis/redis)

## Getting Started

[Docker](https://www.docker.com/get-started) is needed if you want to try the API without setting up your
own database server.

Run the following commands to start experiencing this API:

```shell
# download the starter kit
git clone https://github.com/chaksaray/Golang-REST-API-Boilerplate-with-JWT-and-Mysql.git

cd Golang-REST-API-Boilerplate-with-JWT-and-Mysql

# start a Mysql database server, redis cache & running project in a Docker container
docker-compose up -d

# generate and serve swagger of out api doc
bin/serve-swagger

# test our api
bin/test
```

At this time, you have a REST API server running at `http://127.0.0.1:8080`. It provides the following endpoints:

-   `GET /`: show a welcome page
-   `POST /v1/login`: authenticates a user and generates a JWT
-   `GET /v1/users`: returns a paginated list of the albums
-   `GET /v1/users/:id`: returns the detailed information of an user
-   `POST /v1/users`: creates a new user
-   `PUT /v1/users/:id`: updates an existing user
-   `DELETE /v1/users/:id`: deletes a user

Also provide the endpoints of CRUD posts.

## Project Layout

This API uses the following project layout:

```
.
├── app
│   ├── auth             authentication feature
│   ├── cache            redis cache is implemented here
│   ├── controllers      all logic are here
│   ├── middlewares      middleware func
│   ├── models           all entities and all functions for calling to database
│   ├── startup          where database, route, logger is started
│   ├── utils            for the reusable funtions
|   └── app.go           where the app is run
├── bin                  useful commands
├── config               database & redis cache configuration
├── docs                 api doc, swagger.yml
├── logs                 contain error & info log
├── mysql
|   ├── seeds            seeding database data
│   └── init-db.sql      create database
├── tests
|   ├── integration      integration test
│   └── unit             unit test
├── .env                 environment variables
└── main.go              where our api is started
```
