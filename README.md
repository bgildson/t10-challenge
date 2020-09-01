# T10-CHALLENGE

[![Test Status](https://github.com/bgildson/t10-challenge/workflows/Test%20and%20Send%20Coverage%20Report/badge.svg)](https://github.com/bgildson/t10-challenge/actions?workflow=test)
[![Coverage Status](https://coveralls.io/repos/github/bgildson/t10-challenge/badge.svg?branch=master)](https://coveralls.io/github/bgildson/t10-challenge?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/bgildson/t10-challenge)](https://goreportcard.com/report/github.com/bgildson/t10-challenge)

This repository contains the solution to the [T10 Challenge](./challenge).

## Running the solution

_To follow the steps bellow, you should have installed [Docker](https://docs.docker.com/get-docker/) and [docker-compose](https://docs.docker.com/compose/install/)._

To run locally as production, execute the command bellow

```sh
docker-compose -f docker-compose-prod.yml up --build
```

To execute the bellow database operations, set the **DATABASE_URL** envvar

```sh
# default prod database string connection, change this based on your settings
export DATABASE_URL=postgres://postgres:postgres@localhost:5432/t10_challenge?sslmode=disable
```

Apply the database migrations

```sh
docker run --rm -v $(pwd)/migrations:/migrations --network host migrate/migrate:v4.11.0 -path=/migrations -database ${DATABASE_URL} -verbose up
```

Populate the database with some users

```sh
docker run --rm -v $(pwd)/users.sql:/tmp/users.sql --network host postgres:12-alpine psql -Atx ${DATABASE_URL} -f "/tmp/users.sql"
```

The file "[t10-challenge.postman_collection.json](./t10-challenge.postman_collection.json)" contains a **Postman Collection** to interact with the challenge solution.
