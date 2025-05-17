.PHONY: build 
build:
	docker compose build

.PHONY: up
up:
	docker compose up

.PHONY: clean
clean:
	docker compose down --remove-orphans --volumes
