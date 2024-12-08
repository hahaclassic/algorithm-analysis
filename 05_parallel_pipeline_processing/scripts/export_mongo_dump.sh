#!/bin/bash

if [ -z "$MONGO_INITDB_ROOT_USERNAME" ] || [ -z "$MONGO_INITDB_ROOT_PASSWORD" ] || [ -z "$MONGO_INITDB_DATABASE" ] || [ -z "$MONGO_INITDB_HOST" ] || [ -z "$MONGO_INITDB_PORT" ]; then
    echo "env variables not found"
    exit 1
fi

MONGO_URI="mongodb://${MONGO_INITDB_ROOT_USERNAME}:${MONGO_INITDB_ROOT_PASSWORD}@${MONGO_INITDB_HOST}:${MONGO_INITDB_PORT}"

echo "export data..."
mongoexport --authenticationDatabase=$MONGO_INITDB_ROOT_USERNAME \
            --uri="${MONGO_URI}" \
            --db="${MONGO_INITDB_DATABASE}" \
            --collection="recipes" \
            --out="/jsondump/recipes.json" \
            --jsonArray --pretty

echo "Done. dir with data: ../jsondump/recipes.json"
