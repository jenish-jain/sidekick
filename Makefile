.PHONY: all
export GO111MODULE=on

APP=sidekick
APP_EXECUTABLE="$(APP)"
IMAGE_TAG="$(shell git describe)"
ALL_PACKAGES=$(shell go list ./... | grep -v "vendor")

help:
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done


compile:
	go build -o $(APP_EXECUTABLE) -ldflags "-X 'github.com/jenish-jain/sidekick/internal/version.version=$(IMAGE_TAG)'" *.go

fmt:
	go fmt $(ALL_PACKAGES)

build: # complies format vet and lint your code
build: compile fmt

run: # run's your go code from source
run:
	go run cmd/main.go

start: # start's your generated binary
start:
	${APP_EXECUTABLE} start

tar: build
	tar -czf sidekick.tar.gz sidekick README.md

clean:
	rm -rf out/
