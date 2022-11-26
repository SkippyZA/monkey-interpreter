SHELL=/bin/bash
.DEFAULT_GOAL := help

.PHONY: help test test/cov test/cov-html

## test: Run unit tests
test:
	@go test -v ./...

## test/cov: Run unit tests and generate a coverage report
test/cov:
	@mkdir -p .coverage
	@set -o pipefail; go test -v ./... -covermode='count' -coverprofile=.coverage/coverage.out | tee .coverage/test-coverage.out

## test/cov-html: Open the HTML code coverage report
test/cov-html:
	go tool cover -html=.coverage/coverage.out


## :
## help: Print out available make targets.
help: Makefile
	@echo
	@echo " Choose a command run:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
