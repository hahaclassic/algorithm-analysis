include ./code/.env
export $(shell sed 's/=.*//g' ./code/.env)

startdb:
	docker compose up mongodb

dump:
	docker cp ./scripts/export_mongo_dump.sh mongodb:/export_mongo_dump.sh
	docker exec -it mongodb chmod +x /export_mongo_dump.sh
	docker exec -it mongodb bash -c "/export_mongo_dump.sh"

gantt:
	python3 ./scripts/gantt_chart.py

.PHONY: dump startdb