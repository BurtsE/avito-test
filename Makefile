
.PHONY: run ## Start application (requires running Compose services)
run:
	@echo Starting application...
	@dotenv -f ./.env run -- env ${dev-env-vars} go run ./cmd