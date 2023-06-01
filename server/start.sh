#!/bin/sh

set -e

echo "Initializing application"

make dropdb
make createdb
make migrateup
make build
make dev