.PHONY: all clean deps build

all: deps test-unit build

test-unit:
	@echo ">> running unit tests"
	@go test -short -coverprofile=unit.coverprofile -covermode=atomic -race ./...
