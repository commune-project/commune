#!/bin/bash

go test github.com/commune-project/commune/models || exit $?
go test github.com/commune-project/commune/db/dbmanagers || exit $?
go test github.com/commune-project/commune/ap/asgenerator || exit $?

echo "All tests are passed!"