
.PHONY: run ## Start application 
run:
	@echo Starting application...
	@dotenv -f ./.env run -- env ${dev-env-vars} go run ./cmd

.PHONY: runf## Start filler
runf:
	@echo Starting application...
	@dotenv -f ./.env run -- env ${dev-env-vars} go run ./filler

.PHONY: dockerup## Start application (requires running Compose services)
dockerup:
	@echo Starting application...
	@dotenv -f ./.env run -- env ${dev-env-vars} docker-compose  -f deploy/compose.yml up