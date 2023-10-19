.PHONY: run
run:
	$(GOENV) go run cmd/app/main.go $(RUN_ARGS)

.PHONY: migrate
migrate:
	$(GOENV) go run cmd/migrate/main.go $(RUN_ARGS)
