# How to Perform Database Migrations using Go Migrate

## Migrate CLI Installation
```shell
go install -tags postgres github.com/golang-migrate/migrate/v4/cmd/migrate@v4.16.2
go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@v4.16.2
```
```shell
vim ~/.profile
```
```
export PATH=$PATH:/data/code/go/bin/
```
```shell
source ~/.profile
```

## Create a new migration file

create a directory to store all the migration files.
```shell
mkdir migrations
cd migrations
```

create migration files using the following command:
```shell
migrate create -ext sql -dir ./migrations -seq ums_init_schema
migrate create -ext sql -dir ./migrations -seq cms_init_schema
migrate create -ext sql -dir ./migrations -seq discovery_init_schema
```
`-seq` to generate a sequential version and `init_user_schema` is the name of the migration.

A migration typically consists of two distinct files, one for moving the database to a new state (referred to as "up") and another for reverting the changes made to the previous state (referred to as "down").

## Fill migration files
Next you need to populate the file using the appropriate SQL query.

## Run migration
```
go run cmd/migrate/main.go
```
