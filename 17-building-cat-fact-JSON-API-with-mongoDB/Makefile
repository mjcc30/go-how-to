EXEC = building_cat_fact_json_api_with_mongoDB
BIN_SRC=bin/$(EXEC)

.PHONY: init
init:
	@go mod init $(EXEC)

.PHONY: build
build:
	@go build -o $(BIN_SRC)

.PHONY: run
run: build
	./$(BIN_SRC)

.PHONY: test
test:
	@go test -v ./...
