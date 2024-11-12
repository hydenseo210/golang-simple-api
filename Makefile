build:
	docker build --build-arg LASTCOMMITSHA="testsha" -t hyden-simple-api .

.PHONY: test
test: lint ## Run tests and create coverage report
	@test -z "$(BUILDKITE)" || echo "~~~ :golang: unit tests"
	go test -race -short -coverprofile=coverage.txt -covermode=atomic ./... \
		&& go tool cover -func=coverage.txt

.PHONY: lint
lint: ## Lint Go code
	@test -z "$(BUILDKITE)" || echo "~~~ :golang: linting go"
	go vet ./...

run:
	docker run -i -t -d --platform linux/amd64 -p=8080:8080 --name="hyden-simple-api" hyden-simple-api 

up:
	build run

stop:
	docker stop hyden-simple-api
	docker rm hyden-simple-api

clean:
	docker stop hyden-simple-api
	docker rm hyden-simple-api
	docker image rm hyden-simple-api