# README #

## Setup

1. Install Go version 1.20
2. Install Mockery version v2.13 or later
3. Use GoLand (recommended)
4. Download dependencies with command `go mod download` or `go mod tidy`
5. Create `.env` file based on `.env.example`
6. Create database "myDate" or anything you want, you can set up the connection to database on db/gorm.go
7. Run Migration with command `make migrate_up` then will be created schema on database 

## Run

Use this command to run API app from root directory:

```shell
go run cmd/api/main.go
```

## Unit Tests

### Generate Mocks

To generate mock, run:

```shell
mockery --all --keeptree --case underscore --with-expecter
```

### Run Unit Tests

To run unit tests:
```shell
go test ./...
```

---
