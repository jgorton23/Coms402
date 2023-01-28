include .env.example
export

# HELP =================================================================================================================
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

# compose-up: ### Run docker-compose
# 	docker-compose up --build -d postgres && docker-compose logs -f
# .PHONY: compose-up

# compose-up-integration-test: ### Run docker-compose with integration test
# 	docker-compose up --build --abort-on-container-exit --exit-code-from integration
# .PHONY: compose-up-integration-test

# compose-down: ### Down docker-compose
# 	docker-compose down --remove-orphans
# .PHONY: compose-down

# run: #swag-v1 
# 	go mod tidy && go mod download && \
# 	CGO_ENABLED=0 go run ./cmd/app
# .PHONY: run

# docker-rm-volume: ### remove docker volume
# 	docker volume rm go-clean-template_pg-data
# .PHONY: docker-rm-volume

# linter-golangci: ### check by golangci linter
# 	golangci-lint run
# .PHONY: linter-golangci

# linter-hadolint: ### check by hadolint linter
# 	git ls-files --exclude='Dockerfile*' --ignored | xargs hadolint
# .PHONY: linter-hadolint

# linter-dotenv: ### check by dotenv linter
# 	dotenv-linter
# .PHONY: linter-dotenv

# linter-gci: ## sort imports
# 	gci write . --skip-generated -s standard,default
# .PHONY: linter-imports

test: ### run test
	go test -v -cover -race ./internal/...
.PHONY: test

integration-test: ### run integration-test
	go clean -testcache && go test -v ./integration-test/...
.PHONY: integration-test

mock: ### run mockery
	mockery --all -r --case snake
.PHONY: mock

gen-db: ### gen database from models 
	go run -mod=mod entgo.io/ent/cmd/ent generate ./pkg/database/models --target ./pkg/database/ent
	make linter-imports
.PHONY: gen-db

gen-api-spec: ### gen http openapi spec 
	cd api/v1 && npm run build
.PHONY: gen-api-spec

gen-http:  gen-api-spec ### gen http interface from openapi spec 
	oapi-codegen --config oapi-codegen-config.yaml api/v1/_build/openapi.yaml > internal/delivery/controller/v1.gen.go
.PHONY: gen-http

# TODO convert to command
# go run -mod=mod entgo.io/ent/cmd/ent init --target ./pkg/database/models {User}