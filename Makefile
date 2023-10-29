.PHONY: docker-up
docker-up:
	@docker-compose up --build -d

.PHONY: docker-down
docker-down:
	@docker-compose down

.PHONY: run-app
run-app:
	@make -s docker-up
	@sh ./docker/check-net.sh 5432
	@go mod vendor
	@go run -race -v main.go