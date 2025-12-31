# Identify the path to the main.go file
MAIN_PACKAGE_PATH := ./cmd/api

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## run/dev: run the cmd/api application in development mode
.PHONY: run dev
run dev:
	go run ${MAIN_PACKAGE_PATH}

## run/prod: run the cmd/api application in production mode
.PHONY: run/prod
run/prod:
	go run ${MAIN_PACKAGE_PATH} -port=4000 -env=production

# ==================================================================================== #
# BUILD
# ==================================================================================== #

