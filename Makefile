.DEFAULT_GOAL := dev

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: lint
	go vet ./...
	# shadow ./...
.PHONY:vet

dev: vet
	@reflex -r '.go' -s -- go run cmd/main.go

postgres:
	docker run --rm -it --name postgresWallet -p 5434:5432 -e POSTGRES_USER=root  -e POSTGRES_PASSWORD=e4dd99ae701 -d postgres:13.6-alpine
.PHONY:postgres

createdb:
	docker exec -it postgresWallet createdb --username=root --owner=root wallet-engine-db
.PHONY:createdb

elastic:
	docker run --rm -it -d  -p 9200:9200 -p 9300:9300  -e "discovery.type=single-node"  docker.elastic.co/elasticsearch/elasticsearch:8.1.0
.PHONY:elastic

rabbitmq:
	docker run --rm -it -d -p 15672:15672 -p 5672:5672 --hostname my-rabbit --name evea-rabbit -e RABBITMQ_DEFAULT_USER=root -e RABBITMQ_DEFAULT_PASS=e4dd99ae701 rabbitmq:3-management
.PHONY:rabbitmq

seed:
	go run seeder/2022022201-seed.go
.PHONY:seed

test:
	go test -v -cover ./...
.PHONY:seed

start: |
	go install github.com/swaggo/swag/cmd/swag@v1.8.1
	 ./start.sh
.PHONY:swagger


