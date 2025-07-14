.PHONY : help, test, dev-run, build, cover-html, run, clean

_DEFAULT := help

test:
	@echo "preparing tests with their % coverage"
	@go test -cover ./...

run: build run

build:
	@go build -o rubic

dev-run:
	@go run main.go

cover-html:
	@go test -coverprofile=coverage.out
	@go tool cover -html=coverage.out -o coverage.html

clean:
	@rm -rf coverage.* 
