SHELL:=/bin/bash
.ONESHELL:

# go test -v ./...
test:
	set -euo pipefail
	go clean -testcache && \
	go test -v . | sed ''/PASS/s//$$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$$(printf "\033[31mFAIL\033[0m")/''
.PHONY:test

format:
	go fmt .


run:
	go run main.go 29cb773e13f2e7a5ab4271d7772762ccae0ea057