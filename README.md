# Go Sample API using Echo framework
___

## Features
- Basic CRUD API utilizing sqlite database
- TTL Cache for database query results
- Open API Spec/Swagger endpoint

## TODO
- initialize DB + migrations support
- stub external api call
- Unit + Integration tests

## Requirements
- Install Go `https://golang.org/doc/install`
- Table must exist in sqlite DB prior to using API. Create statement defined: [000001_create_items_table.up.sql](db/migrations/000001_create_items_table.up.sql)
  - TODO included to remove this manual step

## Environment Variables
Required for DB connection \
Defined: [config.go](config/config.go)

## Running the server
Defined in [Makefile](Makefile)

Run
> make run

Build binary
> make build

## Swagger
Update docs within `/docs` directory: `make build-swagger`

Run server and navigate to url: `localhost:1323/swagger/`
