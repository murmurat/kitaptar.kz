# Tools.
TOOLS = tools
TOOLS_BIN = $(TOOLS)/bin


protoc:
	protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative ./protobuf/service.proto

fix-lint:
	golangci-lint run --fix

go-imports:
	goimports -local "kitaptar.kz" -w ./internal ./cmd


#.PHONY: fix-lint
#fix-lint: $(TOOLS_BIN)/golangci-lint
#	$(TOOLS_BIN)/golangci-lint run --fix


#imports: $(TOOLS_BIN)/goimports
#	$(TOOLS_BIN)/goimports -local "kitaptar.kz" -w ./internal ./cmd


# INSTALL linter
$(TOOLS_BIN)/golangci-lint: export GOPATH = C:/Users/meiro/go/$(TOOLS_BIN)
$(TOOLS_BIN)/golangci-lint:
	#mkdir -p $(TOOLS_BIN)
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.2


# INSTALL goimports
$(TOOLS_BIN)/goimports: export GOPATH = C:/Users/meiro/go/$(TOOLS_BIN)
$(TOOLS_BIN)/goimports:
	go install golang.org/x/tools/cmd/goimports@latest



migrate-up:
	migrate -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -path migrations/pg -verbose up 2


migrate-down:
	migrate -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -path migrations/pg -verbose down

migration-create:
	migrate create -ext sql -dir migrations/pg/ -seq

migration-version:
	migrate -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -path migrations/pg force 2

swag-init:
	swag init --parseDependency -g handler/router.go

go-test:
	go test ./... -coverprofile=coverage.out