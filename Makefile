.PHONY: docker-up
docker-up:
	@docker-compose up --build -d

.PHONY: docker-down
docker-down:
	@docker-compose down
