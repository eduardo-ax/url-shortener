.PHONY: build
build: 
	go build -o build/urlshortener

.PHONY: test
test:
	go test -v  ./...

.PHONY: run
run:
	go run main.go

.PHONY: docker_compose_up
docker_compose_up:
	docker compose -p urlshortener up -d	