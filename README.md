# Street Market Microservice

## Building

`docker build -t street-market .`

## MongoDB

`docker run -d -p 27017:27017 mongo`

## Running

### Locally

`go run main.go [api] [--settings]`

`go run main.go [workers] [import-data] [--path] [--settings]`

#### Test Coverage
`go test -failfast -timeout 30s -cover ./...`

### Docker

`docker run street-market [api] [--settings]`

`docker run street-market [workers] [import-data] [--path] [--settings]`

## Flags

* `--settings` path to config.yaml config file (default "config.yaml")
* `--path` path to import the csv file

## Swagger

`http://localhost:9001/swagger`

`https://swagger.io/tools/swagger-ui/`

