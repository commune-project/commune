#!/bin/bash

env PGPASSWORD=defaultPassword dropdb --username=postgres -h postgres commune_go_test
env PGPASSWORD=defaultPassword createdb --username="postgres" -h postgres commune_go_test

sh test_all.sh

echo "All tests are passed!"

env PGPASSWORD=defaultPassword dropdb --username=postgres -h postgres commune_go_test