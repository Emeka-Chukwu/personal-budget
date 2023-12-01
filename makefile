DB_URL=postgresql://admin:password@localhost:5432/personal-budget?sslmode=disable

network:
	docker network create bank-network


createdb:
	docker exec -it postgres createdb --username=admin --owner=password personal-budget

dropdb:
	docker exec -it postgres dropdb personal-budget

migrateup:
	migrate -path migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path migration -database "$(DB_URL)" -verbose down 1

new_migration:
	migrate create -ext sql -dir migration -seq $(name)

test:
	go test -v -cover -short ./...

server:
	go run main.go

# mock:
# 	mockgen -package mockdb -destination db/mock/store.go github.com/techschool/simplebank/db/sqlc Store
# 	mockgen -package mockwk -destination worker/mock/distributor.go github.com/techschool/simplebank/worker TaskDistributor



.PHONY: network postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 new_migration test server mock