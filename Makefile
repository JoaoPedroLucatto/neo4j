.PHONY: build 
build:
	docker compose build

.PHONY: up
up:
	docker compose up

.PHONY: clean
clean:
	docker compose down --remove-orphans --volumes

.PHONY: logs
logs:
	docker compose logs -f --tail=100
