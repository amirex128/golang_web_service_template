export ROOT=$(realpath $(dir $(firstword $(MAKEFILE_LIST))))
export BIN=$(ROOT)/bin
export GOBIN?=$(BIN)
export GOFLAGS=-mod=vendor
export GO=$(shell which go)
export BUILD=cd $(ROOT) && $(GO) install -v -ldflags "-s"

export SELLOORA_MYSQL_MAIN_DB="selloora"
export SELLOORA_MYSQL_MAIN_HOST="localhost"
export SELLOORA_MYSQL_MAIN_PASSWORD="a6766581"
export SELLOORA_MYSQL_MAIN_PORT="3306"
export SELLOORA_MYSQL_MAIN_USER="selloora"
export SELLOORA_REDIS_DB=1
export SELLOORA_REDIS_HOST="localhost"
export SELLOORA_REDIS_PORT=6379
export SERVER_HOST="0.0.0.0"
export SERVER_PORT="8585"
export SERVER_URL="http://localhost:8585"

all:
	$(BUILD) ./cmd/...

convey:
	$(GO) get -v github.com/smartystreets/goconvey
	$(GO) install -v github.com/smartystreets/goconvey

test: convey
	cd $(ROOT)/services && $(GO) test -v -race ./...

test-gui: convey
	cd $(ROOT)/services && goconvey -host=0.0.0.0


migration: tools-go-bindata
	export GOFLAGS=-mod= && $(GO) get -v github.com/mattn/go-sqlite3/...
	export GOFLAGS=-mod= && $(GO) get -v github.com/rubenv/sql-migrate/...
	export GOFLAGS=-mod= && $(GO) install -v github.com/rubenv/sql-migrate/...
	rm -rf $(ROOT)/modules/models/migrations/*.gen.go
	export GOFLAGS=-mod= && cd $(ROOT)/modules/models/migrations && $(BIN)/go-bindata -o $(ROOT)/modules/models/migrations/migration.gen.go -nomemcopy=true -pkg=migrations ./db/...


migcreate:
	@/bin/bash $(ROOT)/scripts/create_migration.sh
mig-up:
	cd $(BIN) && ./migration --action=up

mig-down:
	$cd $(BIN) && ./migration --action=down

tools-migration:
	$(BUILD) $(ROOT)/cmd/migration
