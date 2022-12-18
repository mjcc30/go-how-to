EXEC = structure_golang_API_projects
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
