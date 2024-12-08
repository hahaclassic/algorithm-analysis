include .env
export $(shell sed 's/=.*//g' .env)

MONGO_URI=mongodb://${MONGO_INITDB_ROOT_USERNAME}:${MONGO_INITDB_ROOT_PASSWORD}@${MONGO_INITDB_HOST}:${MONGO_INITDB_PORT}/?authSource=admin
MONGO_DB=${MONGO_INITDB_DATABASE}
MONGO_COLLECTION=recipes
OUTPUT_FILE=recipes.json

startdb:
	docker compose up mongodb

dump:
	docker cp ./scripts/export_mongo_dump.sh mongodb:/export_mongo_dump.sh
	docker exec -it mongodb chmod +x /export_mongo_dump.sh
	docker exec -it mongodb bash -c "/export_mongo_dump.sh"

.PHONY: dump startdb