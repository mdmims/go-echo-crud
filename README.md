# Go Sample API using Echo framework
Echo: https://echo.labstack.com/

## TODO
- initialize DB + migrations support
- in-memory TTL cache for DB queries
- stub external api call

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
