.PHONY: run
run:
	$(GOENV) go run cmd/app/main.go $(RUN_ARGS)

.PHONY: migrate
migrate:
	$(GOENV) go run cmd/migrate/main.go $(RUN_ARGS)

.PHONY: migrate-minio
migrate-minio:
	$(GOENV) go run cmd/minio/main.go $(RUN_ARGS)

.PHONY: gen-docs
gen-docs:
	~/go/bin/swag init -g cmd/app/main.go   --parseDependency --parseInternal