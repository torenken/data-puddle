SHELL := /bin/bash

.DEFAULT_GOAL := all
all: build

# =====
# Building tooling

dev-tooling:
	go build -o ./build/datagen ./app/tooling/datagen/main.go
	go build -o ./build/data-export-dl ./app/tooling/data-export-dl/main.go

# ==============================================================================
# Building lambdas

AWS_PROFILE := # configure aws profile here

NAME := data-puddle
LAMBDA_DIR := ./app/services
BUILD_DIR := ./build

# List of aws lambda functions (see services folder)
ALL_LAMBDAS := \
	provide-ticket-data \
	provide-ticket-url

build: $(addprefix build-, $(ALL_LAMBDAS))

build-%:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o ${BUILD_DIR}/$*/bootstrap ${LAMBDA_DIR}/$*


zip-%:
	zip -j ${BUILD_DIR}/$*/bootstrap.zip ${BUILD_DIR}/$*/bootstrap

deploy-%:
	aws lambda update-function-code --function-name ${NAME}-$* --zip-file fileb://${BUILD_DIR}/$*/bootstrap.zip --no-cli-pager --profile ${AWS_PROFILE}

# ==============================================================================
# Running from within aws
$(ALL_LAMBDAS):
	$(MAKE) build-$@
	$(MAKE) zip-$@
	$(MAKE) deploy-$@

# ==============================================================================
# Modules support

tidy:
	go mod tidy

deps-list:
	go list -m -u -mod=readonly all

deps-upgrade:
	go get -u -v ./...
	go mod tidy