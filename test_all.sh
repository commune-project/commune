#!/bin/sh

go run main.go migrate up || exit $?

go test github.com/commune-project/commune/models || exit $?
go test github.com/commune-project/commune/db/dbmanagers || exit $?
go test github.com/commune-project/commune/ap/asgenerator || exit $?
go test github.com/commune-project/commune/ap/inbox || exit $?
go test github.com/commune-project/commune/handlers || exit $?
go test github.com/commune-project/commune/interfaces/validators || exit $?
