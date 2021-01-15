#!/bin/bash

env PGPASSWORD=defaultPassword dropdb --username=postgres -h postgres commune_go_test
env PGPASSWORD=defaultPassword createdb --username="postgres" -h postgres commune_go_test

go run main.go migrate up

go test github.com/commune-project/commune/models || exit $?
go test github.com/commune-project/commune/db/dbmanagers || exit $?
go test github.com/commune-project/commune/ap/asgenerator || exit $?

echo "All tests are passed!"

env PGPASSWORD=defaultPassword dropdb --username=postgres -h postgres commune_go_test