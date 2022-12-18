start:
	docker-compose up

stop:
	docker-compose down --remove-orphans

start-locally:
	go run ./server.go

rebuild:
	@make stop
	docker-compose build
	docker-compose up --force-recreate --build

unit-test:
	go test -v ./pkg/... -coverprofile=coverage.out

coverage: unit-test
	go tool cover -html=coverage.out

integration-test-locally:
	go test -v ./tests/

integration-test:
	@make stop
	@make start &
	while ! echo exit | nc 0.0.0.0 3333; do sleep 1; done
	@make integration-test-locally
	@make stop

linter:
	golangci-lint run

